//Referencias:
// Time Idea from: https://www.golangprograms.com/how-to-get-the-current-date-and-time-with-timestamp-in-local-and-other-timezones.html
//// Time format from: https://gobyexample.com/time-formatting-parsing


package main
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
	"time"
)
type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
	usersChan = make(map[string]net.Conn)
)

type clientData struct {
  user string
  admin bool
}

var listClients []clientData
var ConUsers int = 0

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	who := ""
	isAdmin := ConUsers == 0;
	go clientWriter(conn, ch)
	input := bufio.NewScanner(conn)
	for input.Scan() {
		mssg := input.Text()
		if match, _ := regexp.MatchString("<user>.+", mssg); match {
			newClient := clientData{user: mssg, admin: isAdmin}
			listClients = append(listClients, newClient)
			ConUsers++
			separatedString := strings.Split(mssg, ">")
			who = separatedString[1]
			fmt.Printf("irc-server > Nuevo usuario conectado [%s] \n", who)
			ch <- "irc-server > bienvenido a IRC server"
			ch <- "irc-server > Tu usuario  [" + who + "] se conecto de manera exitosa"
			if isAdmin {
				ch <- "irc-server > Felicidades eres el primer usuario"
				ch <- "irc-server > Eres el nuevo Administrador"
				fmt.Printf("irc-server > [%s] Fue promovido a Adminsitrador\n", who)
			}
			messages <- "irc-server > Nuevo usuario coenctado [" + who + "]"
			entering <- ch
			usersChan[who] = conn
		} else if mssg == "/users" {
			fmt.Fprintf(conn, "irc-server > ")
			for key := range usersChan {
				fmt.Fprintf(conn, "%s, ", key)
			}
			fmt.Fprintf(conn, "\n")
		} else if match, _ := regexp.MatchString("^/msg .+ .+", mssg); match {
			strimString := strings.Split(mssg, " ")
			lenSlice := len(strimString)
			if user, check := usersChan[strimString[1]]; check {
				fmt.Fprintf(user, "%s [privateMSG] > ", who)
				for i := 2; i < lenSlice; i++ {
					fmt.Fprintf(user, " %s ", strimString[i])
				}
				fmt.Fprintf(user, "\n")
			} else {
				fmt.Fprintf(conn, "irc-server > No user [%s] found, use /users to see connected users. \n", strimString[1])
			}
		} else if mssg == "/time" {
			loc, _ := time.LoadLocation("America/Mexico_city")
    	tme := time.Now().In(loc).Format("15:04")
    	fmt.Fprintf(conn, "irc-server > Local Time: %s %s\n", loc, tme)
		} else if match, _ := regexp.MatchString("^/user .+$", mssg); match {
			strimString := strings.Split(mssg, " ")
			if user, check := usersChan[strimString[1]]; check {
				fmt.Fprintf(conn, "irc-server > username: %s, IP: %s \n", strimString[1], user.RemoteAddr().String())
			} else {
				fmt.Fprintf(conn, "irc-server > No user (%s) found, use /users to see connected users. \n", strimString[1])
			}
		} else if match, _ := regexp.MatchString("^/kick .+$", mssg); match {
			strimString := strings.Split(mssg, " ")
			user := who
			kickUser := strimString[1]
			if listClients[0].user == user {
				if listClients[0].admin {
					fmt.Fprintf(conn, "irc-server > Kickeaste al usuario [%s]\n", kickUser)
					fmt.Printf("irc-server > [%s] fue kickeado\n", kickUser)
				} else {
					fmt.Fprintf(conn, "irc-server > Fuiste kickeado de este canal\n")
					fmt.Fprintf(conn, "irc-server > No se permite el mal uso del lenguaje  BAN-HAMMER\n")
					usersChan[strimString[1]].Close()
					delete(usersChan, kickUser)
				}
				messages <- "irc-server > [" + kickUser + "] El Admin se enfermo de poder y te correo"
			} else {
					ch <- "irc-server > Solo el administrador del canal puede hacer est accion"
			}
		} else {
			messages <- who + "> " + input.Text()
		}
	}
	leaving <- ch
	messages <- "irc-server > [" + who + "] se fue"
	fmt.Printf("irc-server > [%s] abandono \n", who)
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
func main() {
	host := flag.String("host", "localhost", "host string")
	port := flag.String("port", "9000", "port string")
	flag.Parse()
	fmt.Printf("irc-server > El servidor inicio a %s:%s \n", *host, *port)
	fmt.Printf("irc-server > Listo para recibir nuevos clientes \n")
	listener, err := net.Listen("tcp", *host+":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

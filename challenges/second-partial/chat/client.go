
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strings"
)


func main() {
	user := flag.String("user", "defaultUser", "User Name String")
	server := flag.String("server", "localhost:9000", "Host address String")
	flag.Parse()
	conn, err := net.Dial("tcp", *server)
	if err != nil {
		log.Fatal(err)
	}
	if blank := strings.TrimSpace(*user) == ""; blank {
		log.Panic("No puede ser vacio el usuario")
	}
	io.WriteString(conn, "<user>"+*user+"\n")
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Desconectado")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//Diego Gracia
//Eduardo Alonso Herrera A01361404

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	//Con este Codigo le puedo dar color a la salida en consola
	COLOROBJETOS = "\033[38;5;%dm%s\033[39;49m\n"
)

func actPantalla() {
	print("\033[H\033[2J")
	mapaGUI()
	tablaDatos()
	fmt.Printf("\033[%d;0H", tOffset+5)
}

func actTodo() {
	for {
		<-time.After(100 * time.Millisecond)
		go actPantalla()
	}
}

var caracteristicaMapa [][][]int
var pelotas = [][]int{}
var estados = []string{}

//Info pelotas
var numPelotas = 0
var iterations = 0
var alturaMaxima = 0
var n = 0

//Num de pelotas
var norte = 0
var oeste = 0
var sur = 0
var este = 0
var isla = 0

var tOffset int

func entradas() {
	fmt.Println("Tamano del mapa (nxn):")
	_, err := fmt.Scan(&n)
	if err != nil {
		n = 5
	}
	fmt.Println("Altura Maxima:")
	_, err = fmt.Scan(&alturaMaxima)
	if err != nil {
		alturaMaxima = 10
	}
	fmt.Println("Iteraciones:")
	_, err = fmt.Scan(&iterations)
	if err != nil {
		iterations = 1
	}
	fmt.Println("Pelotas por iteracion:")
	_, err = fmt.Scan(&numPelotas)
	if err != nil {
		numPelotas = 1
	}
}

func mapaGUI() {
	for i := 0; i < len(caracteristicaMapa); i++ {
		for j := 0; j < len(caracteristicaMapa); j++ {
			fmt.Printf("\033[%d;%dH", i+3, (j*3)+70)
			if caracteristicaMapa[i][j][1] == 0 {
				//OBJETOS cambia el numero indicado para dar el verde
				fmt.Printf(COLOROBJETOS, 2, "X")
			} else {
				//OBJETOS con el 1 es rojo
				fmt.Printf(COLOROBJETOS, 1, "o")
			}
		}
	}
}

func generarMapa() [][][]int {
	mapaCreado := [][][]int{}
	for i := 0; i < n; i++ {
		row := [][]int{}
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				s := []int{0, 0}
				row = append(row, s)
			} else {
				s := []int{rand.Intn(alturaMaxima), 0}
				row = append(row, s)
			}
		}
		mapaCreado = append(mapaCreado, row)
	}
	return mapaCreado
}

//Esta funcion  crea las corrutinas para cada pelota  cada iteracion crea numPelotas Goroutines.
func incLluvia() {
	var wg sync.WaitGroup
	k := 0
	for i := 0; i < iterations; i++ {
		for j := 0; j < numPelotas; j++ {
			wg.Add(1)
			go lluvia(k, &wg)
			k++
		}
		//Se pone de uno para que sea rapido el proceso
		time.Sleep(time.Second * 3)
	}
	wg.Wait()
	actPantalla()
}

//Esta funciion mneja el comportamiento de las pelotas cuando esta funcion termina llama a sync.WaitGroup, que es el mismo que se usa en incLluvia()
func lluvia(index int, wg *sync.WaitGroup) {
	x, y := posPelotas(index)
	caida(index, x, y)

	for true {
		pelotas[index][1] = y
		pelotas[index][2] = x
		pelotas[index][3] = caracteristicaMapa[x][y][0]
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		if x == 0 || y == 0 || x == (n-1) || y == (n-1) {
			estados[index] = seaPelotas(index, x, y)
			caracteristicaMapa[x][y][1] = 0
			break
		} else if y < (n-1) && caracteristicaMapa[x][y+1][0] < caracteristicaMapa[x][y][0] && caracteristicaMapa[x][y+1][1] == 0 {
			caracteristicaMapa[x][y][1] = 0
			caracteristicaMapa[x][y+1][1] = 1
			velocidad(index, caracteristicaMapa[x][y+1][0], "s")
			y += 1
		} else if x < (n-1) && caracteristicaMapa[x+1][y][0] < caracteristicaMapa[x][y][0] && caracteristicaMapa[x+1][y][1] == 0 {
			caracteristicaMapa[x][y][1] = 0
			caracteristicaMapa[x+1][y][1] = 1
			velocidad(index, caracteristicaMapa[x+1][y][0], "e")
			x += 1
		} else if y > 0 && caracteristicaMapa[x][y-1][0] < caracteristicaMapa[x][y][0] && caracteristicaMapa[x][y-1][1] == 0 {
			caracteristicaMapa[x][y][1] = 0
			caracteristicaMapa[x][y-1][1] = 1
			velocidad(index, caracteristicaMapa[x][y-1][0], "n")
			y -= 1
		} else if x > 0 && caracteristicaMapa[x-1][y][0] < caracteristicaMapa[x][y][0] && caracteristicaMapa[x-1][y][1] == 0 {
			caracteristicaMapa[x][y][1] = 0
			caracteristicaMapa[x-1][y][1] = 1
			velocidad(index, caracteristicaMapa[x-1][y][0], "w")
			x -= 1
		} else {
			estados[index] = "Estanque"
			isla += 1
			break
		}

	}
	pelotas[index][1] = y
	pelotas[index][2] = x
	velocidad(index, caracteristicaMapa[x][y][0], "_")
	wg.Done()
}

func posPelotas(index int) (int, int) {
	x := 1 + rand.Intn(n-2)
	y := 1 + rand.Intn(n-2)
	for caracteristicaMapa[x][y][1] == 1 {
		x = 1 + rand.Intn(n-2)
		y = 1 + rand.Intn(n-2)
	}
	caracteristicaMapa[x][y][1] = 1
	pelotas[index][1] = x
	pelotas[index][2] = y
	return x, y
}

//Controla lacaida de la velocidad de la caida de las pelotas.
func caida(index int, x int, y int) {
	estados[index] = "Caida"
	t := 500
	totalTime := 0
	for pelotas[index][3] >= caracteristicaMapa[x][y][0] {
		pelotas[index][3] = int(pelotas[index][6]*(totalTime/1000)) + pelotas[index][3]
		pelotas[index][6] = pelotas[index][6] + int(-9.8*float64(totalTime/1000))
		totalTime += t
		time.Sleep(time.Millisecond * time.Duration(t))
	}
	pelotas[index][3] = caracteristicaMapa[x][y][0]
	estados[index] = "Movimiento"
}

//Hace el tracking de las pelotas que caen al mar y donde quedan.
func seaPelotas(index int, x int, y int) string {
	if x == 0 {
		norte += 1
		return "Oeste"
	}
	if y == 0 {
		oeste += 1
		return "Sur"
	}
	if x == (n - 1) {
		sur += 1
		return "Este"
	}
	if y == (n - 1) {
		este += 1
		return "Norte"
	}
	return ""
}

func velocidad(index int, zPos int, direction string) {
	angle := math.Atan((float64(pelotas[index][3]) / float64(zPos)))
	if direction == "n" {
		pelotas[index][5] = int(float64(pelotas[index][6]) * math.Cos(angle))
		pelotas[index][6] = int(float64(pelotas[index][6]) * math.Sin(angle))
	} else if direction == "w" {
		pelotas[index][4] = int(float64(pelotas[index][6]) * math.Cos(angle))
		pelotas[index][6] = int(float64(pelotas[index][6]) * math.Sin(angle))
	} else if direction == "s" {
		pelotas[index][4] = -int(float64(pelotas[index][6]) * math.Cos(angle))
		pelotas[index][6] = int(float64(pelotas[index][6]) * math.Sin(angle))
	} else if direction == "e" {
		pelotas[index][5] = -int(float64(pelotas[index][6]) * math.Cos(angle))
		pelotas[index][6] = int(float64(pelotas[index][6]) * math.Sin(angle))
	} else {
		pelotas[index][4] = 0
		pelotas[index][5] = 0
		pelotas[index][6] = 0
	}
}

func initPelotas() {
	for i := 0; i < (numPelotas * iterations); i++ {
		ball := []int{i, 0, 0, rand.Intn(alturaMaxima*20) + alturaMaxima, 0, 0, 0}
		pelotas = append(pelotas, ball)
		estados = append(estados, "Espera")
	}
}

func finalOffset() {
	if (numPelotas * iterations) > n {
		tOffset = (numPelotas * iterations)
	} else {
		tOffset = n
	}
}

func tablaDatos() {
	fmt.Printf("\033[2;0H")
	fmt.Printf("PN | PY | PX |  PZ  | VX | VY |  VZ  |  Estado  |")
	fmt.Printf("\033[3;0H")
	for i := 0; i < len(pelotas); i++ {

		fmt.Printf("\033[%d;0H", i+4)
		fmt.Printf("%d", pelotas[i][0])

		fmt.Printf("\033[%d;4H", i+4)
		fmt.Printf("|")

		fmt.Printf("\033[%d;6H", i+4)
		fmt.Printf("%d", pelotas[i][1])

		fmt.Printf("\033[%d;9H", i+4)
		fmt.Printf("|")

		fmt.Printf("\033[%d;11H", i+4)
		fmt.Printf("%d", pelotas[i][2])

		fmt.Printf("\033[%d;14H", i+4)
		fmt.Printf("|")

		fmt.Printf("\033[%d;16H", i+4)
		fmt.Printf("%d", pelotas[i][3])

		fmt.Printf("\033[%d;21H", i+4)
		fmt.Printf("|")

		fmt.Printf("\033[%d;23H", i+4)
		fmt.Printf("%d", pelotas[i][4])

		fmt.Printf("\033[%d;26H", i+4)
		fmt.Printf("|")

		fmt.Printf("\033[%d;28H", i+4)
		fmt.Printf("%d", pelotas[i][5])

		fmt.Printf("\033[%d;31H", i+4)
		fmt.Printf("|")

		fmt.Printf("\033[%d;33H", i+4)
		fmt.Printf("%d", pelotas[i][6]*-1)

		fmt.Printf("\033[%d;38H", i+4)
		fmt.Printf("|")

		fmt.Printf("\033[%d;40H", i+4)
		fmt.Printf(estados[i])

		fmt.Printf("\033[%d;48H", i+4)
		fmt.Printf("|")
	}
	fmt.Printf("\033[%d;50H", 2)
	fmt.Printf("Isla:")
	fmt.Printf("\033[%d;58H", 2)
	fmt.Printf("%d", isla)

	fmt.Printf("\033[%d;50H", 3)
	fmt.Printf("Oeste:")
	fmt.Printf("\033[%d;58H", 3)
	fmt.Printf("%d", norte)

	fmt.Printf("\033[%d;50H", 4)
	fmt.Printf("Sur:")
	fmt.Printf("\033[%d;58H", 4)
	fmt.Printf("%d", oeste)

	fmt.Printf("\033[%d;50H", 5)
	fmt.Printf("Este:")
	fmt.Printf("\033[%d;58H", 5)
	fmt.Printf("%d", sur)

	fmt.Printf("\033[%d;50H", 6)
	fmt.Printf("Norte:")
	fmt.Printf("\033[%d;58H", 6)
	fmt.Printf("%d", este)
}

//Funcion main
func main() {
	entradas()
	print("\033[H\033[2J")
	//Semilla para generar un numero aleatoreo
	rand.Seed(time.Now().UnixNano())
	caracteristicaMapa = generarMapa()
	initPelotas()
	finalOffset()
	mapaGUI()
	time.Sleep(time.Second)
	go actTodo()
	incLluvia()
}

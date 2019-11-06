package main

import (
	"os"
	"time"
	"fmt"
	"strconv"
)

var i int
var seconds int

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		i = 0
		seconds := 0
		fmt.Println("Number of communications per second:")
	loop:
		for {
			ch1 <- struct{}{}
			select {
			case <-ch2:
				i++
			case <-ticker.C:
				seconds++
				<-ch2 // Must consume ch2, otherwise dead lock
				i++
				fmt.Printf("\rSecond: %d: %d", seconds, 2*i/seconds)
				if seconds >= 6 { //you can change the senconds manually
					ticker.Stop()
					break loop
				}
			}
		}

		ch3 <- struct{}{}
	}()

	go func() {
		for {
			<-ch1
			ch2 <- struct{}{}
		}
	}()
	<-ch3
	fmt.Println()

	file, err := os.Create("exercise9-5.txt")
  if err != nil {
    fmt.Println(err)
    return
  }
  str := "Messages sent: "+strconv.Itoa(i)
  fmt.Println(str)
  out, err := file.WriteString(str)
  if err != nil {
    fmt.Println(err)
    file.Close()
    return
  }
  fmt.Println(out)
}


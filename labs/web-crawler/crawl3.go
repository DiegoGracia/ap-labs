package main

import (
	"fmt"
	"log"
	"flag"
	"github.com/adonovan/gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	for _, elem := range list {
		fmt.Println(elem)
	}
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist
	var depth int;
	var depthcounter int = 0;

	// Start with the command-line arguments.
	n++
	flag.IntVar(&depth, "depth", 3, "User defined depth")
	flag.Parse()
	go func () { worklist <- flag.Args() }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if depthcounter < depth {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					depthcounter++
					n++
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}
}

//!-

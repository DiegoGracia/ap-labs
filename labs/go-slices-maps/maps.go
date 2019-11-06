package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _,c := range strings.Fields(s){
		m[string(c)]+=1;
	}
	return m
}

func main() {
	wc.Test(WordCount)
}


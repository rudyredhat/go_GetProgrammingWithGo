// gophers assembly line
// This gopher doesn’t read values, but only sends them.
package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	downstream <- ""
}

// filters bad from the assembly line
func filterGopher(upstream, downstream chan string) {
	for {
		item := <-upstream
		for item == "" {
			downstream <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

// prints at the end of the assembly
func printGopher(upsteram chan string) {
	for {
		v := <-upsteram
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}
func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}

// hello world
// goodbye all

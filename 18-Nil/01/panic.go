// nil is a zero val
// nil leads to panic if not handled with if
package main

import "fmt"

func main() {
	var nowhere *int
	fmt.Println(nowhere)
	// fmt.Println(*nowhere)       // will cause panic , so comment this
	if nowhere != nil {
		fmt.Println(*nowhere)
	}

}

// <nil>
// panic: runtime error: invalid memory address or nil pointer dereference

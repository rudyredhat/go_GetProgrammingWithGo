// use of fallthrough keyword
package main

import "fmt"

func main() {
	var t1 = "text"
	switch t1 {
	case "text":
		fmt.Println("Got it")
		fallthrough
	case "hello":
		fmt.Println("Hello there")
	default:
		fmt.Println("Will not print this")

	}
}

// Got it
// Hello there

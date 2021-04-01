package main

import "fmt"

func main() {
	var t1 = "text"
	switch t1 {
	case "test":
		fmt.Println("Test abcd")
		// list of possible comma separated val
	case "text", "hello":
		fmt.Println("Text abcd")
	case "outer":
		fmt.Println("Outer abcd")
	default:
		fmt.Println("Bye abcd")

	}
}

// Text abcd

// basic if eg
package main

import "fmt"

func main() {
	var cmd = "hello"
	if cmd == "hello" {
		fmt.Println("Hi hru")
	} else if cmd == "bye" {
		fmt.Println("Get lost")
	} else {
		fmt.Println("Sorry")
	}
}

// Hi hru

// not operator
package main

import "fmt"

func main() {
	// make sure to pass true and false as a boolean value
	var t1 = false
	var t2 = true
	// if any of the statement is true, its true, or logical operator
	if !t1 || !t2 {
		fmt.Println("Nothing to do ")
	}
}

// Nothing to do

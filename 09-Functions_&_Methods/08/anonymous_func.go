// An anonymous function, also called a function literal in Go, is a function without a name.
// You can assign an anonymous function to a variable and then use that variable like
//  any other function, as shown in the following listing.
package main

import "fmt"

var f = func() {
	fmt.Println("Hello Boy")
}

func main() {
	f()
}

// Hello Boy

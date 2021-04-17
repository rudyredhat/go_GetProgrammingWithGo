// slice and array declaration
package main

import "fmt"

func main() {
	arr := [...]string{"asd", "dasd", "dasd", "sfd", "xcv"}

	// slice declare
	sl := []string{"asd", "czxc", "frd"}

	fmt.Printf("array %T\n", arr)
	fmt.Printf("slice %T", sl)
}

// array [5]string
// slice []string

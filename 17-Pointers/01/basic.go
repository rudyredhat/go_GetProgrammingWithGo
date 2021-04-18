// reference and dereferece operator
// check pointer type
// declare a pointer
package main

import "fmt"

func main() {
	ans := 42
	fmt.Println(&ans) // location in the memory

	// dereference operator
	add := &ans
	fmt.Println(*add)

	// pointer type
	fmt.Printf("address is a %T\n", add)

	// declare a pointer
	canada := "Canada"
	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &canada
	fmt.Println(*home)
}

// 0xc000012088
// 42
// address is a *int
// home is a *string
// Canada

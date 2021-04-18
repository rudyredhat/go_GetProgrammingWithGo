// pointers with array
package main

import "fmt"

func main() {
	superpower := &[3]string{"flight", "invisibility", "super strength"}
	//  array in the previous listing is dereferenced automatically when indexing or slicing it.
	fmt.Println(superpower[0])
	fmt.Println(superpower[1:2])
}

// flight
// [invisibility]

// When a variable is declared to be an interface type without an assignment, the zero value is nil
package main

import "fmt"

func main() {
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v)

	// var of the interface is defined
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	//%#v format verb is shorthand to see both type and value,
	fmt.Printf("%#v\n", v)
}

// <nil> <nil> <nil>
// *int <nil> false
// (*int)(nil)

// Interfaces are concerned with what a type can do, not the value it holds
//   interfaces are declared with a set of methods that a type must satisfy.
// so interfaces provides polymorphism - many shapes
package main

import (
	"fmt"
	"strings"
)

// declare a var within an interface type
// The variable t can hold any value of any type that satisfies the interface.
var t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(3))
}

func main() {
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}

// nack nack
// pew pew pew

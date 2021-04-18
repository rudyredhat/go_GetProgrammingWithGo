// Instead of using a pointer, one alternative is to declare a small structure with a few methods.
// It requires a little more code, but it doesn’t require a pointer or nil.
package main

import "fmt"

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}
func main() {
	n := newNumber(42)
	fmt.Println(n)

	e := number{}
	fmt.Println(e)
}

// 42
// not set

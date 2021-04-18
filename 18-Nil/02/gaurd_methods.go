// guarding methods with nil
// calling nil receiver
package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	//p.age++ // pointer derefrence err - remove this line to run err free
	// calling a nil receiver
	if p == nil {
		return
	}
	p.age++
}
func main() {
	var nobody *person
	fmt.Println(nobody)
	nobody.birthday()
}

// <nil>
// panic: runtime error: invalid memory address or nil pointer dereference

// Method receivers are similar to parameters.
package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++
}
func main() {
	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry)

	//method call without pointer
	nathan := person{
		name: "Nathan",
		age:  17,
	}
	nathan.birthday()
	fmt.Printf("%+v\n", nathan)
}

// &{name:Terry age:16}
// {name:Nathan age:18}

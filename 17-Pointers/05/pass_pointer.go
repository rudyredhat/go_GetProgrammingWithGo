// enable mutation
// pass pointers as a parameter
package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)
}

// {name:Rebecca superpower:imagination age:15}

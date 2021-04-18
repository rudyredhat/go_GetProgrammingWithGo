// pointers with structures
package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	timmy := &person{
		name: "Rudy",
		age:  10,
	}
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
}

// &{name:Rudy superpower:flying age:10}

// declaration of array
// composite literals
package main

import "fmt"

func main() {
	var p [10]string
	p[0] = "hello" // assign a value
	p[1] = "are"
	p[2] = "you"

	// access the value
	loc1 := p[0]
	fmt.Println("At p[o]", loc1)

	// len of p
	fmt.Println("len of p: ", len(p))
	fmt.Println("At p[3]", p[3] == "")
}

// At p[o] hello
// len of p:  10
// At p[3] true

//----------------------composite literals
// dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
// planets := [...]string{      1
// "Mercury",
// "Venus",
// "Earth",
// "Mars",
// "Jupiter",
// "Saturn",
// "Uranus",
// }

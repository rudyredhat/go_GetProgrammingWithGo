// use of pass by value
// no change in the main function and the values residing in the array
// arrays not used as function parameters
package main

import "fmt"

func terraform(planets [8]string) {
	// The terraform function is operating on a copy of the planets array,
	// so the modifications don’t affect planets in the main function.
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terraform(planets)
	fmt.Println(planets)
}

// [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]

//------------------
/*

dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
terraform(dwarfs)

Can’t use dwarfs (type [5]string) as type [8]string in argument to terraform

For these reasons, arrays aren’t used as function parameters nearly as often as slices.


*/

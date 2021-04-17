// use of append function
package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Earth")
	fmt.Print(dwarfs)
}

// [Ceres Pluto Haumea Makemake Eris Earth]

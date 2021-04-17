// avoid extra allocations and copies by preallocating a slice with the built-in make function.
// specifies both the length (0) and capacity (10)
// Up to 10 elements can be appended before dwarfs runs out of capacity, causing append to allocate a new array.
package main

import "fmt"

func lencap(label string, slice []string) {
	fmt.Printf("%v, length: %v, capacity: %v, %v\n", label, len(slice), cap(slice), slice)
}
func main() {
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	lencap("dwarfs1: ", dwarfs)
	dwarfs = append(dwarfs, "Ceres1", "Pluto1", "Haumea1", "Makemake1", "Eris1")
	lencap("dwarfs1: ", dwarfs)
	dwarfs = append(dwarfs, "Ceres2", "Pluto2", "Haumea2", "Makemake2", "Eris2")
	lencap("dwarfs1: ", dwarfs)
}

// dwarfs1: , length: 5, capacity: 10, [Ceres Pluto Haumea Makemake Eris]
// dwarfs1: , length: 10, capacity: 10, [Ceres Pluto Haumea Makemake Eris Ceres1 Pluto1 Haumea1 Makemake1 Eris1]
// dwarfs1: , length: 15, capacity: 20, [Ceres Pluto Haumea Makemake Eris Ceres1 Pluto1 Haumea1 Makemake1 Eris1 Ceres2 Pluto2 Haumea2 Makemake2 Eris2]

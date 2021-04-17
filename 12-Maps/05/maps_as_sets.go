// A set is a collection similar to an array, except that each element is guaranteed to occur only once
// sort the final set
package main

import (
	"fmt"
	"sort"
)

func main() {
	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool) // makes a map with boolean values
	for _, t := range temperatures {
		set[t] = true
	}
	if set[-28.0] {
		fmt.Println("set member")
	}
	fmt.Println(set)
	// but the keys printed above are always in arbitary order
	// sort them
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
		fmt.Println(unique)
	}

	sort.Float64s(unique)
	fmt.Println("Final after sort")
	fmt.Println(unique)
}

// map[-33:true -31:true -29:true -28:true -23:true 32:true]
// [-23]
// [-23 -33]
// [-23 -33 -28]
// [-23 -33 -28 32]
// [-23 -33 -28 32 -31]
// [-23 -33 -28 32 -31 -29]
// Final after sort
// [-33 -31 -29 -28 -23 32]

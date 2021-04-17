// group the set of temp in maps and slices
package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64) // A map with float64 keys and []float64 value

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
		fmt.Println(groups)
		fmt.Println(groups[g])
	}
	fmt.Println()
	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}

}

// map[-20:[-28]]
// [-28]
// map[-20:[-28] 30:[32]]
// [32]
// map[-30:[-31] -20:[-28] 30:[32]]
// [-31]
// map[-30:[-31] -20:[-28 -29] 30:[32]]
// [-28 -29]
// map[-30:[-31] -20:[-28 -29 -23] 30:[32]]
// [-28 -29 -23]
// map[-30:[-31] -20:[-28 -29 -23 -29] 30:[32]]
// [-28 -29 -23 -29]
// map[-30:[-31] -20:[-28 -29 -23 -29 -28] 30:[32]]
// [-28 -29 -23 -29 -28]
// map[-30:[-31 -33] -20:[-28 -29 -23 -29 -28] 30:[32]]
// [-31 -33]

// -20: [-28 -29 -23 -29 -28]
// 30: [32]
// -30: [-31 -33]

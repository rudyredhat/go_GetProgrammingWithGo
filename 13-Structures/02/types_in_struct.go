// use of types in structure to use the same fields
package main

import "fmt"

func main() {
	type location struct { // type struct
		lat  float64
		long float64
	}
	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734
	fmt.Println(spirit, opportunity)
}

// {-14.5684 175.472636} {-1.9462 354.4734}

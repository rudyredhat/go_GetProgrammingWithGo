// create a single slice from the struct and configure each as single unit
package main

import "fmt"

func main() {
	type location struct {
		name      string
		lat, long float64
	}
	locations := []location{
		{name: "Brad", lat: -21.231, long: 313.31},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}
	fmt.Println(locations)
}

// [{Brad -21.231 313.31} {Columbia Memorial Station -14.5684 175.472636} {Challenger Memorial Station -1.9462 354.4734}]

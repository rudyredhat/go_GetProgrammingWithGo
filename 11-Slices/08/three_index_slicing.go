// three-index slicing
// check the diff with len and cap in slice
package main

import "fmt"

func lencap(label string, slice []string) {
	fmt.Printf("%v, length: %v, capacity: %v, %v\n", label, len(slice), cap(slice), slice)
}
func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets[0:4:4] // if not defined it will take 8 as default
	lencap("terestrial: ", terrestrial)
	worlds := append(terrestrial, "Ceres")
	lencap("worlds: ", worlds)
	lencap("planets: ", planets)

}

// terestrial: , length: 4, capacity: 4, [Mercury Venus Earth Mars]
// worlds: , length: 5, capacity: 8, [Mercury Venus Earth Mars Ceres]
// planets: , length: 8, capacity: 8, [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]

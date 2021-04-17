// copy the slice to a var - which changes the underlying values of other slices
package main

import "fmt"

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

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice)

	// assign slice to a new var
	iceGiantsMarkII := iceGiants
	// change the previous slice
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiants, iceGiantsMarkII, ice)
}

// [Mercury Venus Earth Mars] [Jupiter Saturn] [Uranus Neptune]
// [Jupiter Saturn Uranus Neptune] [Jupiter Saturn] [Uranus Neptune]
// [Mercury Venus Earth Mars Jupiter Saturn Uranus Poseidon]
// [Uranus Poseidon] [Uranus Poseidon] [Uranus Poseidon]

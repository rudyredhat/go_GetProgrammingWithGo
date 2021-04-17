// If you need a composite literal that’s anything
// more than a list of values, consider writing a constructor function.
package main

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64 // degree, minute and sec intialize
	h       rune    // alias of int32
}

type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 { // method
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// construct a new location
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {

	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal())
	//curiosity := location{lat.decimal(), long.decimal()}
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)
}

// -4.5895 137.4417
// {-4.5895 137.4417}

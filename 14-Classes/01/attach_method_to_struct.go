// attaching methods to structures
package main

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64 // degree, minute and sec intialize
	h       rune    // alias of int32
}

func main() {

	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal())
}
func (c coordinate) decimal() float64 { // method
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// -4.5895 137.4417

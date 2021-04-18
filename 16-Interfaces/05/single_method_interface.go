// single-method interfaces
// Go encourages composition over inheritance, using simple, often one-method interfaces ...
// that serve as clean, comprehensible boundaries between components.
package main

import "fmt"

type Stringer interface {
	String() string
}

type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	Curiosity := location{-4.5895, 137.4417}
	fmt.Println(Curiosity)
}

// -4.5895, 137.4417

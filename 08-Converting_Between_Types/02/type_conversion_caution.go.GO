// type conversion from float to int
package main

import (
	"fmt"
	"math"
)

func main() {
	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println(h) // will result in null and the program stops
	// so to handle it we can have math package provide Min and Max constants
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		// handle out of range
	}
}

// error - null

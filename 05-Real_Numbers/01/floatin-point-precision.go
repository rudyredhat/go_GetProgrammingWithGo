// use of single and double floating point precision
// use of Printf
package main

import (
	"fmt"
	"math"
)

func main() {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Println(pi64)
	fmt.Println(pi32)
	// use of Printf for number of exact values after floating point
	// use \n if want to use Printf in next statement
	fmt.Printf("%.3f\n", pi64)
	fmt.Printf("%4.2f\n", pi64) //total width of 4 with 2 precision
	// zero padding
	fmt.Printf("%05.2f", pi64)
}

// 3.141592653589793
// 3.1415927
// 3.142
// 3.14
// 03.14

// creating custom types
package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius converts K to C
func kelvinToCelsius(k kelvin) celsius {
	// type must be converted before return
	return celsius(k - 273.15)
}
func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Print(k, " degree K is ", c, " degree C")
}

// 294 degree K is 20.850000000000023 degree C

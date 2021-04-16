// method and function diff
// declaring method
// method calling
// method use case - every type of temperature can have a celsius method to convert to Celsius.
package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// kelvinToCelsius converts K to C
func kelvinToCelsius(k kelvin) celsius {
	// type must be converted before return
	return celsius(k - 273.15)

}

// declaring method - celsius method on the kelvin type
func (k kelvin) celsius() celsius { // func + receiver + method_name + result
	// method must have exactly one receiver
	return celsius(k - 273.15)
}

// each temp can provide a celsius type
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main() {
	var k kelvin = 294.0
	// calling function
	c := kelvinToCelsius(k)
	fmt.Println(k, " degree K is ", c, " degree C")
	// calling method
	c = k.celsius()
	fmt.Println(k, " degree K is ", c, " degree C")
	// callinf fahrenheit method
	c = k.celsius()
	fmt.Print("fah ", c)
}

// 294 degree K is 20.850000000000023 degree C
// 294 degree K is 20.850000000000023 degree C
// fah 20.850000000000023

/*
func Println(a ...interface{}) (n int, err error)
for this: Println is said to be a variadic function.
The parameter a is a collection of the arguments passed to the function.
The type of the a parameter is interface{}, known as the empty interface type.

// 1st eg function declaration = Kelvin to Celsius
*/

package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
func main() {
	Kelvin := 294.0
	Celsius := kelvinToCelsius(Kelvin)
	fmt.Print(Kelvin, " degree K is ", Celsius, " degree C")
}

// 294 degree K is 20.850000000000023 degree C

// interchangeable sensor function
// assigning functions to var
//  Function and method calls always have parentheses, such as fakeSensor(), which isn’t the case here.

package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func realSensor() kelvin {
	return 0
}
func main() {
	// sensor() will effectively call either realSensor or fakeSensor,
	// depending on which function sensor is assigned to.
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}

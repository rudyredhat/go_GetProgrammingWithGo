// passing a func to other func
// declare function types
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

// receiving a function
func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v K\n", k)
		// pause for a sec after every o/p
		time.Sleep(time.Second)
	}
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func main() {
	// passing function inside a function
	measureTemperature(3, fakeSensor)
}

// 156 K
// 206 K
// 193 K

/*
// declare func types

type sensor func() kelvin

func measureTemperature(samples int, s func() kelvin)

can now be written like this:

func measureTemperature(samples int, s sensor)
*/

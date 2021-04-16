// closure keeps a reference to surrounding variables rather than a copy of their values, changes to those
// variables are reflected in calls to the anonymous function.
package main

import "fmt"

type kelvin float64

func main() {
	var k kelvin = 294.0

	sensor := func() kelvin {
		return k
	}

	fmt.Println(sensor())
	k++
	fmt.Println(sensor())
}

// 294
// 295

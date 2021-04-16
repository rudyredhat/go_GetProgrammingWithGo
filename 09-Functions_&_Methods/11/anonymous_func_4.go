// it’s possible for a function to return an existing named function,
// declaring and returning a new anonymous function is more useful
package main

import "fmt"

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}

// 5

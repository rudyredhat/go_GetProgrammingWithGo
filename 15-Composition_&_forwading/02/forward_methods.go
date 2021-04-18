// method forwading in the form of struct embedding - specify a type without field name
// embed structure in line 9,10
package main

import "fmt"

type report struct {
	sol int
	temperature
	location
}
type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}
type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}
func main() {
	report := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}

	fmt.Printf("average %v° C\n", report.average())
	// access  the same data of other struct
	fmt.Printf("%v° C\n", report.high)
	report.high = 32
	fmt.Printf("%v° C\n", report.temperature.high)
}

// average -39.5° C
// -1° C
// 32° C

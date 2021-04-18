// composition with struct - embedding to forward methods
// struct inside struct
package main

import "fmt"

// mix of disparate data
// type report struct {
// 	sol       int
// 	high, low float64
// 	lat, long float64
// }

// so use composition to group related fields
type report struct {
	sol         int
	temperature temperature // temp struct
	location    location    // location struct
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

// cal avg temp, use method
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t,

		location: bradbury}

	fmt.Printf("%+v\n", report)

	fmt.Printf("a balmy %v° C\n", report.temperature.high)

	fmt.Printf("average %v° C\n", t.average())

	fmt.Printf("average %v° C\n", report.temperature.average())

}

// {sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5895 long:137.4417}}
// a balmy -1° C
// average -39.5° C
// average -39.5° C

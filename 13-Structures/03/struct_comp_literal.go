// initialize structures with composite literals
// with field name and val
// only val
// use of + sign to print the field name
// struct copy - changes the orig val
package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	//  variables are initialized using field-value pairs
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	// doesn’t specify field names. Instead, a value must be provided
	// for each field in the same order in which they’re listed in the structure
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)

	// we can use + sign to print the field name
	fmt.Printf("%+v\n", spirit)

	spirit1 := spirit
	spirit1.long += 1.0016
	fmt.Println(spirit, spirit1)
}

// {-1.9462 354.4734}
// {4.5 135.9}
// {-14.5684 175.472636}
// {lat:-14.5684 long:175.472636}
// {-14.5684 175.472636} {-14.5684 176.474236}

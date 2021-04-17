// count maps and slice - best eg for count number of occurence in a slice
package main

import "fmt"

func main() {
	// declare a slice
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	//declare a map with make function
	frequency := make(map[float64]int)

	// iterate through the slice (index, val)
	for _, t := range temperatures {
		frequency[t]++
		// iterate through each freqeuncy temp count
		fmt.Println(frequency[t])
		// prints the map with temp:freq
		fmt.Println(frequency)
	}

	// iterate through the map (key , val)
	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d time \n", t, num)
	}

}

// order of the output changes in every new run
// map[-28:1]
// 1
// map[-28:1 32:1]
// 1
// map[-31:1 -28:1 32:1]
// 1
// map[-31:1 -29:1 -28:1 32:1]
// 1
// map[-31:1 -29:1 -28:1 -23:1 32:1]
// 2
// map[-31:1 -29:2 -28:1 -23:1 32:1]
// 2
// map[-31:1 -29:2 -28:2 -23:1 32:1]
// 1
// map[-33:1 -31:1 -29:2 -28:2 -23:1 32:1]
// +32.00 occurs 1 time
// -31.00 occurs 1 time
// -29.00 occurs 2 time
// -23.00 occurs 1 time
// -33.00 occurs 1 time
// -28.00 occurs 2 time

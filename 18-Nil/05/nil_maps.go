// As with slices, a map declared without a composite literal or the make built-in has a value of nil.
// Maps can be read even when nil, as shown in the following listing, though writing to a nil map will panic.
package main

import "fmt"

func main() {
	var soup map[string]int
	fmt.Println(soup == nil)
	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}
	for i, measurement := range soup {
		fmt.Println(i, measurement)
	}

}

// true

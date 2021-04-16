// declaring types to makes code cleaner
// The type keyword declares a new type with a name and an underlying type
// types conversion
package main

import "fmt"

func main() {
	// celsius type is a new numeric type with the same behavior and representation as a float64
	type Celsius float64
	const degree = 20
	var temp Celsius = degree
	// temp += 10
	// so here Celsius type is a unique type , not type alias
	var warmUp float64 = 10
	// to add warmup , must be converted to celsius type
	temp += Celsius(warmUp)
	fmt.Println(temp)
}

// 30

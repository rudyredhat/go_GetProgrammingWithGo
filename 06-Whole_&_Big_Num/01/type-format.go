// %T and [1] use in Printf
package main

import "fmt"

func main() {
	year := 2018
	// if [1] is not used then we need to pass -> year, year
	fmt.Printf("Type is %T & val is %[1]v\n", year)
	// Tyepe is int & val is 2018

}

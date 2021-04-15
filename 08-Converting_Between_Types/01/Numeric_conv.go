// Basic rules for conversion
// 1 - "10" , first conv str to int and then operate
// 2 - 23, 34, 45.66 = x+y/c , error
// Numeric conversion
package main

import "fmt"

func main() {
	age := 41
	marsAge := float64(age)
	marsDays := 687.0
	earth_days := 365.2425
	marsAge = marsAge * earth_days / marsDays
	fmt.Println("I am ", marsAge, "old on Mars")
}

// I am  21.797587336244543 old on Mars

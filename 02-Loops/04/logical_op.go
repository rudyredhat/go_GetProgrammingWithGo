// check leap year or not
// && and || -> logical operator
package main

import "fmt"

func main() {
	fmt.Println("Check the year 2100 is leap year or not.")
	// As with most programming languages, Go uses short-circuit logic.
	// If the first condition is true (the year is evenly divisible by 400),
	// there’s no need to evaluate what follows the || operator, so it is ignored.
	var leap = (2100%400 == 0) || (2100%4 == 0 && 2100%100 != 0)
	if leap {
		fmt.Println("Leap year", leap)
	} else {
		fmt.Println("Not leap year", leap)
	}
}

// Check the year 2100 is leap year or not.
// Not leap year false

// string conversion of rune or byte = string(a)
// digit to string = strconv.Itoa(a)
// number to string = Sprintf
// ASCII to interger (Atoi)
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))

	// CONVERT DIGITS TO STRING
	countdown := 10
	str := "there are " + strconv.Itoa(countdown) + " bikes"
	fmt.Println(str)

	// usinf Sprintf
	str1 := fmt.Sprintf("There are %v bikes.", countdown)
	fmt.Println(str1)

	countdown, err := strconv.Atoi("10")
	if err != nil {
		fmt.Printf("err")
	}
	fmt.Println(countdown)

	/*
		Go uses static typing, once var is declared cannot change its type
		var a = 10
		a = 0.4
		a = fmt.Sprintf("%v",a)

		// results ERR - a stores only int
		// can use numeric conv first
	*/
}

// πάω!
// there are 10 bikes
// There are 10 bikes.
// 10

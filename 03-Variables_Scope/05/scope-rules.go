// random date generation program for mars landing
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2021
	// switch case generating month
	switch month := rand.Intn(10) + 1; month {
	case 2:
		// day generating the date ret
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}

// AD 2021 2 24

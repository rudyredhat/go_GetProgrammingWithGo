// refracting the above code and generating the random date
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2021
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	days := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, days)
}

// AD 2021 6 28

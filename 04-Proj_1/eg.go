// The table should have four columns:

// The spaceline company providing the service
// The duration in days for the trip to Mars (one-way)
// Whether the price covers a return trip
// The price in millions of dollars

// For each ticket, randomly select one of the following spacelines: Space Adventures, SpaceX, or Virgin Galactic.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// step 1 - create a for loop which prints 10 random values
	fmt.Println("Spaceline		Days Trip type	Price")
	fmt.Println("=============================================")
	for count := 1; count <= 10; count++ {
		// used for round or one-way trip
		num := rand.Intn(2) + 1
		// checked for the random spacecraft and used as switch statement
		Spaceline_no := rand.Intn(3) + 1
		switch Spaceline_no {
		case 1:
			Spaceline := "Virgin Galactic"
			if num == 1 {
				fmt.Println(Spaceline, "	23 One-way	$96")
			} else {
				fmt.Println(Spaceline, "	39 Round-trip	$37")
			}
		case 2:
			Spaceline := "SpaceX		"
			if num == 1 {
				fmt.Println(Spaceline, "	31 One-way	$41")
			} else {
				fmt.Println(Spaceline, "	41 Round-trip	$72")
			}
		default:
			Spaceline := "Space Adventures"
			if num == 1 {
				fmt.Println(Spaceline, "	22 One-way	$50")
			} else {
				fmt.Println(Spaceline, "	22 Round-trip	$100")
			}
		}
	}

}

// Spaceline		Days Trip type	Price
// =============================================
// Virgin Galactic 	39 Round-trip	$37
// Space Adventures 	22 Round-trip	$100
// Virgin Galactic 	39 Round-trip	$37
// Space Adventures 	22 Round-trip	$100
// Virgin Galactic 	23 One-way	$96
// SpaceX		 	31 One-way	$41
// Space Adventures 	22 One-way	$50
// Space Adventures 	22 One-way	$50
// Space Adventures 	22 Round-trip	$100
// Space Adventures 	22 Round-trip	$100

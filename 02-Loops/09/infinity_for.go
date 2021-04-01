// infinity check for
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var degrees = 0
	for {
		fmt.Println(degrees)

		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

// 0-359 then again back to 0 and check rand int = 0
// if not , then again prints 0-359, again same process

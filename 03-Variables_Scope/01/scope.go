// print random 10 numbers
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num, "->", count)
		count++

	}
}

// 2 -> 0
// 8 -> 1
// 8 -> 2
// 10 -> 3
// 2 -> 4
// 9 -> 5
// 6 -> 6
// 1 -> 7
// 7 -> 8
// 1 -> 9

//countdown timer
package main

import (
	"fmt"
	"time"
)

func main() {
	var a = 10
	for a > 0 {
		fmt.Println(a)
		// sleep for 1 sec delay
		time.Sleep(time.Second)
		a--
	}
	fmt.Println("Done")
}

// 10
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// Done

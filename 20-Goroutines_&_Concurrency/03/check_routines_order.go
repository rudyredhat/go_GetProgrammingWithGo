// We can find out which ones finish first by passing an argument to each goroutine.
//Passing an argument to a goroutine is like passing an argument to any function:
// the value is copied and passed as a parameter.
// ** order of goroutine always changes **
//  It’s waiting for four seconds when it only needs to wait for just over three seconds.
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i) // oder will change every time
	}
	time.Sleep(4 * time.Second)
}
func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " snore...")
}

// ... 0  snore...
// ... 3  snore...
// ... 4  snore...
// ... 1  snore...
// ... 2  snore...

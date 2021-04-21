// multiple go routine
// All goroutines appear to run at the same time. They might not technically run at the same time,
// though, because computers only have a limited number of processing units.
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	time.Sleep(4 * time.Second)
	fmt.Println("in main") // at 5th sec prints this
}
func sleepyGopher() {
	fmt.Println("in func")      // 1. prints as soon as the prog runs
	time.Sleep(3 * time.Second) // 4 sec hold ** prob here cover by channels later **
	fmt.Println("...snore...")  // prints this
}

// in func
// in func
// in func
// in func
// in func
// ...snore...
// ...snore...
// ...snore...
// ...snore...
// ...snore...
// in main

// creating channels
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // create a channel
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c // receive val from chnnel
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}
func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore...")
	c <- id // sends a val to main
}

// ...  1  snore...
// ...  2  snore...
// gopher  1  has finished sleeping
// ...  0  snore...
// ...  3  snore...
// gopher  2  has finished sleeping
// gopher  0  has finished sleeping
// ...  4  snore...
// gopher  3  has finished sleeping
// gopher  4  has finished sleeping

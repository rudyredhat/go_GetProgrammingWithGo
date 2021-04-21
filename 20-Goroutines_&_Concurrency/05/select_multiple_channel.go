// creating channels and avoid waiting for them
// deadlocks
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int) // create a channel
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finish sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

// now some gophers need to wake up before time ran out
// can be helpful to limit the amount of time doing something
func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}

// my patience ran out
// --- after adding the below func
// gopher  0  has finish sleeping
// my patience ran out

/*
When one or more goroutines end up blocked for something that can never happen, it’s called deadlock,
and your program will generally crash or hang up. Deadlocks can be caused by something as simple as this:

func main() {
    c := make(chan int)
	<-c
}
*/

// In Go, an independently running task is known as a goroutine.

package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher() // go routine start
	time.Sleep(4 * time.Second)
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}

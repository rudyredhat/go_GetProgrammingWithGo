package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if n := rand.Intn(3); n == 0 {
		fmt.Println("first round")
	} else if n == 1 {
		fmt.Println("second round")
	} else {
		fmt.Println("third round")
	}
}

// third round

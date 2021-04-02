package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch n := rand.Intn(3); n {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("three")
	}
}

// two

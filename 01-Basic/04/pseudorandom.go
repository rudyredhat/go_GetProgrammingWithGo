// example of an off-by-one error
// generate rand number between 1-10
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num1 = rand.Intn(10) + 1
	fmt.Println(num1)
	num1 = rand.Intn(10) + 1
	fmt.Println(num1)
}

// 2
// 8

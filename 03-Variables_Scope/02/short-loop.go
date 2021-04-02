//normal for loop and var declaration
package main

import "fmt"

func main() {
	count := 0
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println("outer", count)
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
// outer 0

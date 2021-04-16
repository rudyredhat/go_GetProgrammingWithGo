// iterate through array using len and range
package main

import "fmt"

func main() {
	item := [5]string{"cat", "dog", "bike", "girl", "boy"}
	// use of len
	for i := 0; i < len(item); i++ {
		item := item[i]
		fmt.Println(i, item)
	}
	fmt.Println("")

	// use of range
	for i, item := range item {
		fmt.Println(i, item)
	}
}

// 0 cat
// 1 dog
// 2 bike
// 3 girl
// 4 boy

// 0 cat
// 1 dog
// 2 bike
// 3 girl
// 4 boy

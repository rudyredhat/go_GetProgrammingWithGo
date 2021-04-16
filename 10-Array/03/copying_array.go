// copying array
package main

import "fmt"

func main() {
	item := [5]string{"cat", "dog", "bike", "girl", "boy"}
	item2 := item

	item[2] = "bmw"

	fmt.Println(item)
	fmt.Println(item2)
}

// [cat dog bmw girl boy]
// [cat dog bike girl boy]

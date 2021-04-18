// A slice that’s declared without a composite literal or the make built-in will have a value of nil.
package main

import "fmt"

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}
func main() {
	var soup []string
	fmt.Println(soup == nil)
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	fmt.Println(len(soup))
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)
	// starting with a nil slice
	soup1 := mirepoix(nil)
	fmt.Println(soup1)
}

// true
// 0
// [onion carrot celery]
// [onion carrot celery]

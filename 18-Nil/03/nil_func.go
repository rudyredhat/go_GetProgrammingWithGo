//  variable is declared as a function type, its value is nil by default
// sort alphabetically
package main

import (
	"fmt"
	"sort"
)

var fn func(a, b int) int

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}
func main() {
	fmt.Println(fn == nil)
	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)
}

// true
// [carrot celery onion]

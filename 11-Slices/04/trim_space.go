// use of strings.TrimSpace and strings.Join using slice
package main

import (
	"fmt"
	"strings"
)

// Both worlds and planets are slices, and though worlds is a copy,
// they both point to the same underlying array.
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func main() {
	planets := []string{"venus  ", "  earth  ", "    jupiter   "}
	hyperspace(planets)

	fmt.Println(planets)
	fmt.Println(strings.Join(planets, ""))
	fmt.Println(strings.Join(planets, " "))
}

// [venus earth jupiter]
// venusearthjupiter
// venus earth jupiter

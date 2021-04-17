package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	// The sort package in the standard library declares a StringSlice type
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}

// [Earth Jupiter Mars Mercury Neptune Saturn Uranus Venus]

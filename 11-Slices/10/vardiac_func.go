package main

import "fmt"

// here we can receive set of strings in worlds - as a vardiac func (printf and append)
func terraform(prefix string, worlds ...string) []string {
	// allocate memory for new worlds - creating a new slice here
	newworlds := make([]string, len(worlds))
	for i := range worlds {
		newworlds[i] = prefix + " " + worlds[i]
	}
	return newworlds
}
func main() {
	// taking "New" as the prefix over here and rest 2 as the set of strings
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	// to pass a slice in the worlds
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)
}

// [New Venus New Mars]
// [New Venus New Mars New Jupiter]

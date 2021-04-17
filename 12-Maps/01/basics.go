// maps basic, keys and values
// comma, ok syntax if values are not in map
package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("Avg Earth temp %v\n", temp)
	temp = temperature["Mars"]
	fmt.Printf("Avg Mars temp %v\n", temp)

	temperature["Earth"] = 45
	temperature["Venus"] = 215
	fmt.Println(temperature)

	// comma, ok syntax
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On moon %v", moon)
	} else {
		fmt.Printf("Where is moon")
	}
}

// Avg Earth temp 15
// Avg Mars temp -65
// map[Earth:45 Mars:-65 Venus:215]

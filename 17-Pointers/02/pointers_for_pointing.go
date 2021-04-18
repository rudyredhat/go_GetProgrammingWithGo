package main

import "fmt"

func main() {
	var admin *string
	a1 := "rudy"
	admin = &a1
	fmt.Println(*admin)

	a2 := "rudra"
	admin = &a2
	fmt.Println(*admin)

	a2 = "abc"
	fmt.Println(*admin)

	*admin = "xyz"
	fmt.Println(*admin)
	fmt.Println(a2)
}

// rudy
// rudra
// abc
// xyz
// xyz

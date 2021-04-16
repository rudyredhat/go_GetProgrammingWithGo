// varibale declared within the main function
package main

import "fmt"

func main() {
	f := func(message string) {
		fmt.Println(message)
	}
	f("Go to the party")
}

// Go to the party

//decalre anonymous function in one step
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Function Anonymous")
	}()
}

// Function Anonymous

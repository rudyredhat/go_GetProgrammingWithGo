// strings package with Contains function
// boolean implementation
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("This is just a example statement.")
	var text1 = "test example"
	var text2 = strings.Contains(text1, "example")
	fmt.Println("I got the ans: ", text2)
}

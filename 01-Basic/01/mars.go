// for win setup = https://www.geeksforgeeks.org/how-to-install-go-on-windows/
// for vs code setup = https://dev.to/ko31/how-to-setup-golang-with-vscode-1i4i#
//:~:text=Installing%20Visual%20Studio%20Code&text=Open%20the%20Extensions%20Marketplace%20
//(%20Cmd,the%20Go%20extensions%20listed%20there.
// debug error = https://blog.golang.org/go116-module-changes
// Weight and Years on Mars ?
// package, import and func = 3 keywords
package main

import "fmt"

func main() {
	fmt.Print("Weight in lbs: ")
	fmt.Println(148 * 0.3783) //mars weight is 38% of earth weight
	fmt.Print("Age: ")
	fmt.Println(41 * 365 / 687) //mars revolve wround sun in 687 days
}

// Weight in lbs: 55.9884
// Age: 21

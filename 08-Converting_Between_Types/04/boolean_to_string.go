// converting a boolean to string
package main

import "fmt"

func main() {
	launch := false
	// boolean to str
	launchText := fmt.Sprintf("%v", launch)
	// then print it
	fmt.Println("Ready to launc:", launchText)
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch", yesNo)
}

// Ready to launc: false
// Ready for launch no

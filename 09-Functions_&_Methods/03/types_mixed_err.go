// types cannot be mixed
package main

import "fmt"

func main() {
	type celsius float64
	type fahrenheit float64
	var c celsius = 20
	var f fahrenheit = 20
	if c == f {
		fmt.Print("Inside if")
	}
	c += f
}

// ERR - (mismatched types celsius and fahrenheit)

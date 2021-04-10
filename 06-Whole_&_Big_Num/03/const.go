// use of const
//  The Go compiler is written in Go. Under the hood,
// untyped numeric constants are backed by the big package, enabling all the usual operations
package main

import "fmt"

func main() {
	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400

	const days = distance / lightSpeed / secondsPerDay
	fmt.Println("Galaxy is ", days, "lights away.")

	// comment below 3 lines to run **
	km := distance                                //  constant 24000000000000000000 overflows int - error
	days := distance / lightSpeed / secondsPerDay // cannot assign to days, decalred const
	fmt.Println(days, "fits into int")
}

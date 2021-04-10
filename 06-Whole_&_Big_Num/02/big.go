// use of big
// var distance int64 = 41.3e12 - use of exponential format in go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	ls := big.NewInt(299792)
	scndsperday := big.NewInt(86400)
	dist := new(big.Int)
	// 24 quintillion. It won’t fit in an int64, so instead you can create a big.Int from a string:
	// The number 24 quintillion is in base 10 (decimal), so the second argument is 10.
	dist.SetString("24000000000000000000", 10)
	fmt.Println("Galaxy is ", dist, "km away.")
	scnds := new(big.Int)
	scnds.Div(dist, ls)

	days := new(big.Int)
	days.Div(scnds, scndsperday)

	fmt.Println("That is ", days, "of travel at light speed.")

}

// Galaxy is  24000000000000000000 km away.
// That is  926568346 of travel at light speed.

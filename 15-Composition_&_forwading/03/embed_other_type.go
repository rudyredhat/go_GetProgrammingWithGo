// other than struct, embed int or float type as well
package main

import "fmt"

type sol int
type report struct {
	sol
	location
	temperature
}
type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}
type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}
func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}
func main() {
	report := report{sol: 15}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
}

// 1431
// 1431

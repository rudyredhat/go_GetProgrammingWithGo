// division and multiplication accuracy in floating point precesion
// use of math.Abs

package main

import "fmt"

func main() {
	// division first
	c1 := 21.0
	fmt.Print((c1/5.0*9.0)+32, "* F\n")
	fmt.Print((9.0/5.0*c1)+32, "* f\n")
	// 69.80000000000001* F
	// 69.80000000000001* f

	// multiplication first
	F1 := (c1 * 9.0 / 5.0) + 32.0
	fmt.Print(F1, "* F\n")
	// 69.8* F

	// floting point in accuracy while addn
	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
	// 0.30000000000000004
	// so its never 0.1+0.2 != 0.3
	// instead use - fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)          1

}

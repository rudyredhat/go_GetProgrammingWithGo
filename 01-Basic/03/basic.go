// mutiple var declaration
// increment and assignment operators
package main

import "fmt"

func main() {
	// muiltple var declrataion
	var weight, age = 45.6, 24
	// assignment op
	weight = weight * 0.467
	fmt.Println("Weight1: ", weight)
	weight *= 0.683
	fmt.Println("Weight2: ", weight)
	// increment op
	age = age + 41
	fmt.Println("Age1: ", age)
	age += 1
	fmt.Println("Age2: ", age)
	age++
	fmt.Println("Age3: ", age)

}

// Weight1:  21.2952
// Weight2:  14.544621600000003
// Age1:  65
// Age2:  66
// Age3:  67

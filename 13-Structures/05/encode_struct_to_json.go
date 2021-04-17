// struct data to json
// exit of error func created
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat, Long float64 // fields must start from Capital letter
	}
	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//{"Lat":-4.5895,"Long":137.4417}

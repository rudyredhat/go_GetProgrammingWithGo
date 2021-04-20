// check if the file exist in current dir else exit the prog with a error val
// basic err check eg
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// basic.go

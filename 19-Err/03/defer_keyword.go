// defer keyword can be used to make sure file is closed propelry
package main

import (
	"fmt"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	// every return statement that follows defer will result in the f.Close() method being called.
	// before return statement is called
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	return err
}
func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
	}
}

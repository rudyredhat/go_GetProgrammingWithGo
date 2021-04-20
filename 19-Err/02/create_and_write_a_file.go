// create a file
package main

import (
	"fmt"
	"os"
)

// create a file
func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	// from here the file is opened, now we can write in the file

	// close the file
	_, err = fmt.Fprintln(f, "Errors are values.")
	// here we checking for any err otherwise will add the next line below
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully")
	f.Close()
	return err
}

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
	}
}

// creates a new file  everytime - proverbs.txt and will add 2 lines and will close the file

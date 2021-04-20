//  languages rely heavily on exceptions for communicating and handling errors. Go doesn’t have exceptions,
// but it does have a similar mechanism called panic.
// for eg - panic("I forgot my towel")
// to recover from panic we have recover ()
package main

import "fmt"

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	panic("I forgot my towel")
}

// I forgot my towel

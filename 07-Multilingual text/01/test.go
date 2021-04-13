// strings basic
//string together characters and place them between quotes, it’s called a literal string
// var str1 := "string-val"
// use of escape seq
package main

import "fmt"

func main() {
	fmt.Println("I need to do something big\nBut i need to work hard.")
	// multi line prints
	fmt.Println(`
	Use this techq
	To print multi line
	`)
}

// I need to do something big
// But i need to work hard.

// 	Use this techq
// 	To print multi line
// ---------------------------------------------------
// An alias is another name for the same type, so rune and int32 are interchangeable.
// type byte = uint8

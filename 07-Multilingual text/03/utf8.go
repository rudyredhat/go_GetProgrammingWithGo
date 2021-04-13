// decode a spanish package
// use of utf8 and decode rune in string
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	fmt.Println(len(question), "bytes")
	// first step to supporting other languages is to decode characters
	// to the rune type before manipulating them.
	fmt.Println(utf8.RuneCountInString(question), "runes")
	// DecodeRuneInString function returns the first character and the number of bytes the character consumed
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First Rune: %c %v bytes", c, size)
}

// 15 bytes
// 12 runes
// First Rune: ¿ 2 bytes

/*
Use of range keyword

for i, c := range question {
    fmt.Printf("%v %c\n", i, c)
}

for _, c := range question { // use of blank identifier
    fmt.Printf("%c ", c)            1
}
*/

// rotate 13 cipher code
package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' { // used to keep spaces and punctuations as it is
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}

// hi international space station

// type interface
package main

import (
	"fmt"
	"strings"
)

// er suffix must be attached
type talker interface {
	talk() string
}

// shout function has a parameter of type talker
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(3))
}

//type crater struct{} --> must implement the talk method

func main() {
	shout(martian{})
	shout(laser(2))
	//shout(crater{})
}

// NACK NACK
// PEW PEW PEW

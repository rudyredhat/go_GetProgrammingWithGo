// embedding with interfaces
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

// embedding works here
type starship struct {
	laser
}

func main() {
	shout(martian{})
	shout(laser(2))
	// satisfies the talker interface
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
}

// NACK NACK
// PEW PEW PEW
// pew pew pew
// PEW PEW PEW

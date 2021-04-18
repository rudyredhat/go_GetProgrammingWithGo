//  used to determine the memory address of a field inside of a structure.
package main

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

// address pointer in go
type character struct {
	name  string
	stats stats
}

func main() {
	player := character{name: "Matthias"}
	levelUp(&player.stats)

	fmt.Printf("%+v\n", player.stats)
}

// {level:1 endurance:56 health:280}

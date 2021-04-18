// implement func over interface
package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

// stardate returns a fictional measure of time.
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

/*
// here the func is limited to earth days
// so implementing above in interface form

func stardate(t time.Time) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 100 + day + h
}
*/

// the stardater interface can be later expanded
type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}
func (s sol) Hour() int {
	return 0
}

func main() {
	day := time.Date(2014, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	// stardate function operates on both sols
	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(s))
}

// 318.2 Curiosity has landed - normal func form
// 1218.2 Curiosity has landed - interface one
// 1086.0 Happy birthday

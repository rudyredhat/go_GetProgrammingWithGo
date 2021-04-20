// create new errs if the parameters are wrong
package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

// grid is a sudoku grid
type Grid [rows][columns]int8

// contains a constructor function that accepts a string for an error message
func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("cut of bounds") // can change and return the main err 28.12 and 28.13 eg
	}
	g[row][column] = digit
	return nil

}

// inBounds function in the next listing ensures that row and column are inside the grid boundaries.
func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

//the main function in the next listing creates a grid and displays any error resulting from an invalid placement.
func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		fmt.Printf("An error occurred: %v.\n", err)
		os.Exit(1)
	}
}

// An error occurred: cut of bounds.

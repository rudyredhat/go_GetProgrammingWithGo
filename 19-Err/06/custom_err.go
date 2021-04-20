// creating new err types
// type assertions covered in 28.17 in the end of this
// Because below prog converts SudokuError to an error interface type before it’s returned,
// you may wonder how to access the individual errors. The answer is with type assertions.
// Using a type assertion, you can convert an interface to the underlying concrete type.
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type error interface {
	Error() string
}

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

const rows, columns = 9, 9

// creating custom error types
type SudokuError []error

// Error returns one or more errors separated by commas.
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

// grid is a sudoku grid
type Grid [rows][columns]int8

// contains a constructor function that accepts a string for an error message
/*func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("cut of bounds") // can change and return the main err 28.11 , 28.12 and 28.13 eg
	}
	g[row][column] = digit
	return nil

}*/

// appending the error
func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
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
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
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

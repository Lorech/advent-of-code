package aoc2025

import "fmt"

// Solves the puzzle for the given day using the provided input.
//
// Layers an abstraction over the main function to allow a simpler way of
// getting the solutions of each given day, without having to implement file
// parsing within the `puzzles` package.
func Solve(day int, input string) (any, any, error) {
	var (
		p1, p2 any
	)
	var err error

	switch day {
	case 1:
		p1, p2 = dayOne(input)
	case 2:
		p1, p2 = dayTwo(input)
	default:
		err = fmt.Errorf("No implemented solution")
	}

	return p1, p2, err
}

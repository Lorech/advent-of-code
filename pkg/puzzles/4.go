package puzzles

import (
	"slices"
	"strings"
)

// Day 4: Ceres Search
// https://adventofcode.com/2024/day/4
func dayFour(input string) (int, int) {
	return d4p1(input), d4p2(input)
}

// Completes the first half of the puzzle for day 4.
func d4p1(input string) int {
	words := 0
	// TODO: Learn how to parse files better. I have to trim it to avoid modifying the unit test :)
	input = strings.TrimSuffix(input, "\n")
	rows := strings.Split(input, "\n")

	for y, row := range rows {
		for x, char := range row {
			// Look up only the start of words.
			if char != 'X' {
				continue
			}

			// Initialize a 7x7 grid of zeroes to perform lookups within.
			grid := make([][]byte, 7)

			// Fill the valid grid postions with the values from the input.
			for gy := range 7 {
				// Allocate memory for the row.
				grid[gy] = make([]byte, 7)

				dy := gy - 3

				// y is out of bounds - leave it as a 0.
				if y+dy < 0 || y+dy > len(rows)-1 {
					continue
				}

				for gx := range 7 {
					dx := gx - 3

					// x is out of bounds - leave it as a 0.
					if x+dx < 0 || x+dx > len(row)-1 {
						continue
					}

					grid[gy][gx] = rows[y+dy][x+dx]
				}
			}

			straight, diagonal := lookup(grid, "XMAS", 0)
			words += straight + diagonal
		}
	}

	return words
}

// Completes the second half of the puzzle for day 4.
func d4p2(input string) int {
	words := 0
	// TODO: Learn how to parse files better. I have to trim it to avoid modifying the unit test :)
	input = strings.TrimSuffix(input, "\n")
	rows := strings.Split(input, "\n")

	for y, row := range rows {
		for x, char := range row {
			// Look up only the center of the grid.
			if char != 'A' {
				continue
			}

			// Initialize a 5x5 grid of zeroes to perform lookups within.
			grid := make([][]byte, 5)

			// Fill the valid grid postions with the values from the input.
			for gy := range 5 {
				// Allocate memory for the row.
				grid[gy] = make([]byte, 5)

				dy := gy - 2

				// y is out of bounds - leave it as a 0.
				if y+dy < 0 || y+dy > len(rows)-1 {
					continue
				}

				for gx := range 5 {
					dx := gx - 2

					// x is out of bounds - leave it as a 0.
					if x+dx < 0 || x+dx > len(row)-1 {
						continue
					}

					grid[gy][gx] = rows[y+dy][x+dx]
				}
			}

			_, diagonals := lookup(grid, "MAS", 1)
			// Include only X-shapes of the keyword.
			if diagonals == 2 {
				words += 1
			}
		}
	}

	return words
}

// Performs omnidirectional lookup for the keyword within a sub-grid.
//
// It is assumed that the sub-grid is an odd-length, square 2D slice, with it's
// center being the keyword symbol at index `kOffset`.
//
// Returns the number of times the keyword was found in the grid - the first
// value represents straight-angle finds (T, R, D, L), while the second one
// represents diagonal finds (TR, BR, BL, TL).
func lookup(grid [][]byte, keyword string, kOffset int) (int, int) {
	center := (len(grid) - 1) / 2
	var straights, diagonals int

	// Allocate memory for storing the word in all 8 directions.
	words := make([][]byte, 8)
	for i := range 8 {
		words[i] = make([]byte, len(keyword))
	}

	// Get the word in all 8 directions around the central letter within the grid.
	directions := [3]int{-1, 0, 1}
	for y := range 3 {
		for x := range 3 {
			// Skip an offset of 0,0, as that will just repeatedly get one letter.
			if x == 1 && y == 1 {
				continue
			}

			for offset := range len(keyword) {
				// Look up a symbol on the grid. We must validate the indices in case
				// we are looking up an offset within the word, not the start of it.
				yPos := center + offset*directions[y] - kOffset*directions[y]
				if yPos < 0 || yPos >= len(grid) {
					break
				}
				xPos := center + offset*directions[x] - kOffset*directions[x]
				if xPos < 0 || xPos >= len(grid) {
					break
				}

				words[y*3+x-B2i(y == 2 || y == 1 && x > 0)][offset] = grid[yPos][xPos]
			}
		}
	}

	// Check how many directions produced a valid keyword.
	for i, word := range words {
		if kOffset != 0 {
			slices.Reverse(word)
		}
		if keyword == string(word) {
			if i == 1 || i == 3 || i == 4 || i == 6 {
				straights++
			} else {
				diagonals++
			}
		}
	}

	return straights, diagonals
}

// Gets an integer representation of a boolean value.
//
// NOTE: This looks like a useful function to extract into a utility.
func B2i(input bool) int {
	if input {
		return 1
	}

	return 0
}

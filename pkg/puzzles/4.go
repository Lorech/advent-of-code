package puzzles

import (
	"strings"
)

// Day 4: Ceres Search
// https://adventofcode.com/2024/day/4
func dayFour(input string) (int, int) {
	return d4p1(input), 0
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

			words += lookup(grid, "XMAS")
		}
	}

	return words
}

// Performs omnidirectional lookup for the keyword within a sub-grid.
//
// It is assumed that the sub-grid is an odd-length, square 2D slice, with its'
// center being the first symbol of the keyword.
//
// Returns the number of times the keyword was found within the grid.
func lookup(grid [][]byte, keyword string) int {
	center := (len(grid) - 1) / 2
	result := 0

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
				words[y*3+x-B2i(y == 2 || y == 1 && x > 0)][offset] = grid[center+offset*directions[y]][center+offset*directions[x]]
			}
		}
	}

	// Check how many directions produced a valid keyword.
	for _, word := range words {
		if keyword == string(word) {
			result++
		}
	}

	return result
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

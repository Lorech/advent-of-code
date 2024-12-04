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
			grid := [7][7]byte{}

			// Fill the valid grid postions with the values from the input.
			for gy := range 7 {
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

			words += lookup(grid)
		}
	}

	return words
}

// Performs omnidirectional lookup for the keyword within a sub-grid.
//
// Based on the requirements of the challenge, it is expected that the grid is
// 7x7 tiles large, centered around the first character of the keyword.
//
// Returns the number of times the keyword was found within the grid.
func lookup(grid [7][7]byte) int {
	result := 0

	// right
	if grid[3][3] == 'X' && grid[3][4] == 'M' && grid[3][5] == 'A' && grid[3][6] == 'S' {
		result++
	}
	// bottom-right
	if grid[3][3] == 'X' && grid[4][4] == 'M' && grid[5][5] == 'A' && grid[6][6] == 'S' {
		result++
	}
	// bottom
	if grid[3][3] == 'X' && grid[4][3] == 'M' && grid[5][3] == 'A' && grid[6][3] == 'S' {
		result++
	}
	// bottom-left
	if grid[3][3] == 'X' && grid[4][2] == 'M' && grid[5][1] == 'A' && grid[6][0] == 'S' {
		result++
	}
	// left
	if grid[3][3] == 'X' && grid[3][2] == 'M' && grid[3][1] == 'A' && grid[3][0] == 'S' {
		result++
	}
	// top-left
	if grid[3][3] == 'X' && grid[2][2] == 'M' && grid[1][1] == 'A' && grid[0][0] == 'S' {
		result++
	}
	// top
	if grid[3][3] == 'X' && grid[2][3] == 'M' && grid[1][3] == 'A' && grid[0][3] == 'S' {
		result++
	}
	// top-right
	if grid[3][3] == 'X' && grid[2][4] == 'M' && grid[1][5] == 'A' && grid[0][6] == 'S' {
		result++
	}

	return result
}

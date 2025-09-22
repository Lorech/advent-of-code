package puzzles

import (
	"slices"
	"strings"
)

// Day 8: Resonant Collinearity
// https://adventofcode.com/2024/day/8
func dayEight(input string) (int, int) {
	return d8p1(input), d8p2(input)
}

// Completes the first half of the puzzle for day 8.
func d8p1(input string) int {
	rows := strings.Split(input, "\n")
	bounds := [2]int{len(rows[0]) - 1, len(rows) - 1}
	antennas := parseAntennas(rows)
	antinodes := make([][2]int, 0)

	for _, coords := range antennas {
		fAntinodes := make([][2]int, 0)

		for a := 0; a < len(coords)-1; a++ {
			for b := a + 1; b < len(coords); b++ {
				ns := findAntinodes(coords[a], coords[b], bounds, 1)
				for _, n := range ns {
					// Make sure there is not already an antinode here.
					if slices.Contains(antinodes, n) {
						continue
					}

					// Exclude antenna coordinates.
					if slices.Contains(coords, n) {
						continue
					}

					// Valid antinode.
					fAntinodes = append(fAntinodes, n)
				}
			}
		}

		// Add all antinodes of this frequency to the total pool.
		antinodes = append(antinodes, fAntinodes...)
	}

	return len(antinodes)
}

// Completes the second half of the puzzle for day 8.
func d8p2(input string) int {
	rows := strings.Split(input, "\n")
	bounds := [2]int{len(rows[0]) - 1, len(rows) - 1}
	antennas := parseAntennas(rows)
	antinodes := make([][2]int, 0)

	for _, coords := range antennas {
		for a := 0; a < len(coords)-1; a++ {
			for b := a + 1; b < len(coords); b++ {
				ns := findAntinodes(coords[a], coords[b], bounds, -1)
				for _, n := range ns {
					// Make sure there is not already an antinode here.
					if slices.Contains(antinodes, n) {
						continue
					}

					// Valid antinode.
					antinodes = append(antinodes, n)
				}
			}
		}
	}

	return len(antinodes)
}

// Returns the antinodes produced by the provided pair of antennas.
//
// The returned slice contains all valid antinodes, taking into consideration
// the provided boundraries (max x and max y coordinates), and the maximum
// number of antinodes that one antenna can produce.
//
// If n < 1, it is assumed that all possible antinodes should be searched.
func findAntinodes(a [2]int, b [2]int, boundraries [2]int, n int) [][2]int {
	antinodes := make([][2]int, 0)
	xd := b[0] - a[0]
	yd := b[1] - a[1]

	// Find all the antinodes from antenna a.
	for steps := 0; true; steps++ {
		x := a[0] - steps*xd
		y := a[1] - steps*yd

		// The node would be out of bounds.
		if x < 0 || x > boundraries[0] || y < 0 || y > boundraries[1] {
			break
		}

		antinodes = append(antinodes, [2]int{x, y})

		// Check if we need more nodes.
		if n > 0 && steps == n {
			break
		}
	}

	// Find all the antinodes from antenna b.
	for steps := 0; true; steps++ {
		x := b[0] + steps*xd
		y := b[1] + steps*yd

		// The node would be out of bounds.
		if x < 0 || x > boundraries[0] || y < 0 || y > boundraries[1] {
			break
		}

		antinodes = append(antinodes, [2]int{x, y})

		// Check if we need more nodes.
		if n > 0 && steps == n {
			break
		}
	}
	return antinodes
}

// Parses the input data into a structure containing antenna location data in
// the form of a map, keyed by the antenna frequency, and a value of a slice
// containing all locations of this antenna in an x,y coordinate int array.
func parseAntennas(rows []string) map[rune][][2]int {
	antennas := make(map[rune][][2]int)

	for y, row := range rows {
		for x, tile := range row {
			// There is nothing here.
			if tile == '.' {
				continue
			}

			// This is an antenna!
			coords := [2]int{x, y}
			_, exists := antennas[tile]
			if exists {
				antennas[tile] = append(antennas[tile], coords)
			} else {
				antennas[tile] = [][2]int{coords}
			}
		}
	}

	return antennas
}

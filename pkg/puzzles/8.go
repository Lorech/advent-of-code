package puzzles

import (
	"slices"
	"strings"
)

// Day 8: Resonant Collinearity
// https://adventofcode.com/2024/day/8
func dayEight(input string) (int, int) {
	return d8p1(input), 0
}

// Completes the first half of the puzzle for day 8.
func d8p1(input string) int {
	rows := strings.Split(input, "\n")
	antennas := parseAntennas(rows)
	antinodes := make([][2]int, 0)

	for _, coords := range antennas {
		for a := 0; a < len(coords)-1; a++ {
			for b := a + 1; b < len(coords); b++ {
				ns := findAntinodes(coords[a], coords[b])
				for _, n := range ns {
					// X is out-of-bounds
					if n[0] < 0 || n[0] >= len(rows[0]) {
						continue
					}

					// Y is out-of-bounds
					if n[1] < 0 || n[1] >= len(rows) {
						continue
					}

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

// Returns the coordinates of the two antinodes that the provided pair of
// antenna coordinates would produce. Always returns both antinodes, so the
// boundrary within the grid must be validated by the caller.
func findAntinodes(a [2]int, b [2]int) [2][2]int {
	xd := b[0] - a[0]
	yd := b[1] - a[1]
	an := [2]int{a[0] - xd, a[1] - yd}
	bn := [2]int{b[0] + xd, b[1] + yd}
	return [2][2]int{an, bn}
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

package puzzles

import (
	"lorech/advent-of-code-2024/pkg/cslices"
	"slices"
	"strings"
)

// Day 23: LAN Party
// https://adventofcode.com/2024/day/23
func dayTwentyThree(input string) (int, int) {
	return d23p1(input), 0
}

// Completes the first half of the puzzle for day 23.
func d23p1(input string) int {
	connections := parseNetwork(input)

	options := make([][]string, 0)
	for root, c := range connections {
		// Only look for connections around a t-node.
		if root[0] != 't' {
			continue
		}

		for i := 0; i < len(c)-1; i++ {
			for j := 1; j < len(c); j++ {
				if slices.Contains(connections[c[i]], c[j]) {
					grid := []string{root, c[i], c[j]}
					slices.Sort(grid)
					if !cslices.ContainsSlice(options, grid) {
						options = append(options, grid)
					}
				}
			}
		}
	}

	return len(options)
}

// Parses the input data into a map, keyed by each computer, with values of
// every computer connected to it.
func parseNetwork(input string) map[string][]string {
	pairs := strings.Split(input, "\n")
	connections := make(map[string][]string)
	for _, connection := range pairs {
		c := strings.Split(connection, "-")
		slices.Sort(c)
		for i, n := range c {
			connections[n] = append(connections[n], c[1-i])
		}
	}
	return connections
}

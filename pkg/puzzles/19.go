package puzzles

import (
	"fmt"
	"slices"
	"strings"
)

// Day 19: Linen Layout
// https://adventofcode.com/2024/day/19
func dayNineteen(input string) (int, int) {
	return d19p1(input), 0
}

// Completes the first half of the puzzle for day 19.
func d19p1(input string) int {
	towels, patterns := parseTowels(input)
	possible := 0

	for _, pattern := range patterns {
		reconstructions := make([]string, 0)

		// Prepare the initial reconstructions based on the first letter.
		options, found := towels[rune(pattern[0])]
		if !found {
			continue
		}

		for _, option := range options {
			if strings.HasPrefix(pattern, option) {
				reconstructions = append(reconstructions, option)
			}
		}

		// Go through all in-progress reconstructions until there are no more and
		// the pattern can't be made, or until we find a match for the pattern.
		for len(reconstructions) > 0 {
			r := reconstructions[len(reconstructions)-1]
			reconstructions = reconstructions[:len(reconstructions)-1]

			// This matches the pattern! No point in further processing.
			if r == pattern {
				possible++
				break
			}

			options, found := towels[rune(pattern[len(r)])]
			if !found {
				continue
			}

			// Persist all of the reconstructions that continue to remain valid.
			for _, option := range options {
				if strings.HasPrefix(pattern[len(r):], option) {
					reconstructions = append(reconstructions, fmt.Sprintf("%s%s", r, option))
				}
			}
		}
	}

	return possible
}

// Parses the input data into structured data:
//   - A map, keyed by the first color of the towel storing all matching towels;
//   - The towel combinations that must be generated.
func parseTowels(input string) (map[rune][]string, []string) {
	rows := strings.Split(input, "\n")
	towels := make(map[rune][]string, 0)
	patterns := make([]string, len(rows)-2)

	for i, row := range rows {
		// Delimiter between towels and patterns
		if len(row) == 0 {
			continue
		}

		r := strings.Split(row, ", ")
		if len(r) > 1 {
			slices.Sort(r)
			for _, towel := range r {
				towels[rune(towel[0])] = append(towels[rune(towel[0])], towel)
			}
		} else {
			patterns[i-2] = r[0]
		}
	}

	return towels, patterns
}

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
		success := createPattern(pattern, towels)
		if success {
			possible++
		}
	}

	return possible
}

// Checks if the given color pattern can be recreated using the provided towel
// colors. Returns true or false depending on if it's possible or not.
func createPattern(pattern string, towels map[rune][]string) bool {
	reconstructions := make([]string, 0)

	// Prepare the initial reconstructions based on the first letter.
	options, found := towels[rune(pattern[0])]
	if !found {
		return false
	}

	for _, option := range options {
		if strings.HasPrefix(pattern, option) {
			reconstructions = append(reconstructions, option)
		}
	}

	// Go through all in-progress reconstructions until there are no more and
	// the pattern can't be made, or until we find a match for the pattern.
	checked := make(map[string][]string, 0)
	for len(reconstructions) > 0 {
		r := reconstructions[len(reconstructions)-1]
		reconstructions = reconstructions[:len(reconstructions)-1]

		// This matches the pattern! No point in further processing.
		if r == pattern {
			return true
		}

		options, found := towels[rune(pattern[len(r)])]
		if !found {
			continue
		}

		// Persist all of the reconstructions that continue to remain valid.
		// Only continue with new combinations to prevent infinite loops where
		// some partials create other partials.
		for _, option := range options {
			if strings.HasPrefix(pattern[len(r):], option) && !slices.Contains(checked[r], option) {
				reconstructions = append(reconstructions, fmt.Sprintf("%s%s", r, option))
				checked[r] = append(checked[r], option)
			}
		}
	}

	return false
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

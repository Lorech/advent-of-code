package puzzles

import (
	"slices"
	"strings"
)

// Day 19: Linen Layout
// https://adventofcode.com/2024/day/19
func dayNineteen(input string) (int, int) {
	return d19p1(input), d19p2(input)
}

// Completes the first half of the puzzle for day 19.
func d19p1(input string) int {
	towels, patterns := parseTowels(input)
	cache := make(map[string]int)
	possible := 0

	for _, pattern := range patterns {
		options := countMatches(pattern, towels, &cache)
		if options > 0 {
			possible++
		}
	}

	return possible
}

// Completes the second half of the puzzle for day 19.
func d19p2(input string) int {
	towels, patterns := parseTowels(input)
	cache := make(map[string]int)
	possible := 0

	for _, pattern := range patterns {
		options := countMatches(pattern, towels, &cache)
		possible += options
	}

	return possible
}

// Counts the amount of unique matches for a given pattern using only the
// provided towels without overlap.
func countMatches(pattern string, towels map[rune][]string, cache *map[string]int) int {
	if len(pattern) == 0 {
		return 0
	}

	matches, found := (*cache)[pattern]
	if found {
		return matches
	}

	options, found := towels[rune(pattern[0])]
	if !found {
		matches = 0
	} else {
		for _, option := range options {
			// If the pattern exactly matches an option, store the number of ways to get it.
			if pattern == option {
				matches++
			}

			// Other, smaller patterns may still be created even if this is already a pattern.
			if strings.HasPrefix(pattern, option) {
				next := pattern[len(option):]
				matches += countMatches(next, towels, cache)
			}
		}
	}

	(*cache)[pattern] = matches
	return matches
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

package puzzles

import (
	"regexp"
	"strconv"
)

// Day 3: Mull It Over
// https://adventofcode.com/2024/day/3
func dayThree(input string) (int, int) {
	return d3p1(input), d3p2(input)
}

// Completes the first half of the puzzle for day 3.
func d3p1(input string) int {
	return multiply(input)
}

// Completes the second half of the puzzle for day 3.
func d3p2(input string) int {
	multiplication := 0
	unparsed := input
	enabled := true

	reEnabled := regexp.MustCompile(`(.|\s)*?don't\(\)`)
	reDisabled := regexp.MustCompile(`(.|\s)*?do\(\)`)

	for unparsed != "" {
		var match string

		if enabled {
			match = reEnabled.FindString(unparsed)
			// If there was no match, the entire remaining string has multiplication enabled.
			if match == "" {
				match = unparsed
			}
			multiplication += multiply(match)
		} else {
			match = reDisabled.FindString(unparsed)
		}

		// If there are no remaining matches, the string is fully processed.
		if match == "" {
			break
		}

		enabled = !enabled
		unparsed = unparsed[len(match):]
	}

	return multiplication
}

// Parses the given memory fragment and multiplies all valid `mul()` calls within.
func multiply(fragment string) int {
	multiplication := 0

	re := regexp.MustCompile(`mul\(\d+?,\d+?\)`)
	matches := re.FindAllString(fragment, -1)

	for _, match := range matches {
		reDigits := regexp.MustCompile(`\d+`)
		digits := reDigits.FindAllString(match, -1)
		first, _ := strconv.Atoi(digits[0])
		second, _ := strconv.Atoi(digits[1])
		multiplication += first * second
	}

	return multiplication
}

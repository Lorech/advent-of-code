package puzzles

import (
	"strconv"
	"strings"
)

// Day 2: Red-Nosed Reports
// https://adventofcode.com/2024/day/2
func dayTwo(input string) (int, int) {
	reports := strings.Split(input, "\n")
	safe := 0

	for _, report := range reports {
		// Parse the report into a slice of integers.
		levels := make([]int, len(strings.Fields(report)))
		for i, field := range strings.Fields(report) {
			levels[i], _ = strconv.Atoi(field)
		}

		// Skip empty reports. TODO: Remove this when input comes in as lines.
		if len(levels) < 1 {
			continue
		}

		// Keep track of orientation for the report.
		incrementing := levels[0] < levels[1]

		// Validate the report.
		for i := 0; i < len(levels); i++ {
			// We made it to the end! We're safe!
			if i == len(levels)-1 {
				safe++
				break
			}

			// Difference must be in range [1,3], otherwise, this is unsafe.
			if difference := levels[i+1] - levels[i]; difference == 0 || max(difference, -difference) > 3 {
				break
			}

			// Every level must be either incrementing or decrementing, otherwise, this is unsafe.
			if incrementing && levels[i] > levels[i+1] || !incrementing && levels[i] < levels[i+1] {
				break
			}
		}
	}

	return safe, 0
}

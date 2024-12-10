package puzzles

import (
	"lorech/advent-of-code-2024/pkg/cslices"
	"strconv"
	"strings"
)

// Day 2: Red-Nosed Reports
// https://adventofcode.com/2024/day/2
func dayTwo(input string) (int, int) {
	return d2p1(input), d2p2(input)
}

// Completes the first half of the puzzle for day 2.
func d2p1(input string) int {
	reports := parseReports(input)
	safe := 0

	for _, report := range reports {
		// NOTE: Using the dampened result to detect if the answer is safe to avoid
		// refactoring validateReport into a split function. Maybe I'll fix it one day.
		_, d := validateReport(report, true)
		if d {
			safe++
		}
	}

	return safe
}

// Completes the second half of the puzzle for day 2.
func d2p2(input string) int {
	reports := parseReports(input)
	safe := 0
	dampened := 0

	for _, report := range reports {
		s, d := validateReport(report, false)
		if s {
			safe++
		} else if d {
			dampened++
		}
	}

	return safe + dampened
}

// Validates a single report.
//
// Returns a tuple of two booleans:
//   - The first boolean indicates whether the report is safe.
//   - The second boolean indicates whether the report is actually safe when
//     taking dampening into account.
func validateReport(levels []int, dampened bool) (bool, bool) {
	// Keep track of orientation for the report.
	incrementing := levels[0] < levels[1]

	for i := 0; i < len(levels); i++ {
		// Skip the last level - since we're comparing with the next level, getting to
		// the last level without an early exit implies that the report is safe.
		if i == len(levels)-1 {
			break
		}

		// Difference must be in range [1,3], otherwise this is unsafe unless dampened.
		if difference := levels[i] - levels[i+1]; difference == 0 || max(difference, -difference) > 3 {
			if dampened {
				return false, false
			} else {
				// FIXME: Brute force approach to validating dampening. I should do better!
				_, prev := validateReport(cslices.Remove(levels, i-1), true)
				_, curr := validateReport(cslices.Remove(levels, i), true)
				_, next := validateReport(cslices.Remove(levels, i+1), true)
				return false, prev || curr || next
			}
		}

		// Every level must be either incrementing or decrementing, otherwise this is unsafe unless dampened.
		if incrementing && levels[i] > levels[i+1] || !incrementing && levels[i] < levels[i+1] {
			if dampened {
				return false, false
			} else {
				// FIXME: Brute force approach to validating dampening. I should do better!
				_, prev := validateReport(cslices.Remove(levels, i-1), true)
				_, curr := validateReport(cslices.Remove(levels, i), true)
				_, next := validateReport(cslices.Remove(levels, i+1), true)
				return false, prev || curr || next
			}
		}
	}

	return !dampened, dampened
}

// Parses the input data into a slice of reports, containing a slice of levels.
func parseReports(input string) [][]int {
	rows := strings.Split(input, "\n")
	reports := make([][]int, len(rows))

	for i, row := range rows {
		levels := make([]int, len(strings.Fields(row)))
		for j, field := range strings.Fields(row) {
			levels[j], _ = strconv.Atoi(field)
		}
		reports[i] = levels
	}

	return reports
}

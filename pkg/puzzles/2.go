package puzzles

import (
	"lorech/advent-of-code-2024/pkg/cslices"
	"strconv"
	"strings"
)

// Day 2: Red-Nosed Reports
// https://adventofcode.com/2024/day/2
func dayTwo(input string) (int, int) {
	reports := strings.Split(input, "\n")
	safe := 0
	dampened := 0

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

		// Validate the report.
		isSafe, isActuallySafe := validateReport(levels, false)
		if isSafe {
			safe++
		} else if isActuallySafe {
			dampened++
		}
	}

	return safe, safe + dampened
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
				_, prev := validateReport(cslices.RemoveInt(levels, i-1), true)
				_, curr := validateReport(cslices.RemoveInt(levels, i), true)
				_, next := validateReport(cslices.RemoveInt(levels, i+1), true)
				return false, prev || curr || next
			}
		}

		// Every level must be either incrementing or decrementing, otherwise this is unsafe unless dampened.
		if incrementing && levels[i] > levels[i+1] || !incrementing && levels[i] < levels[i+1] {
			if dampened {
				return false, false
			} else {
				// FIXME: Brute force approach to validating dampening. I should do better!
				_, prev := validateReport(cslices.RemoveInt(levels, i-1), true)
				_, curr := validateReport(cslices.RemoveInt(levels, i), true)
				_, next := validateReport(cslices.RemoveInt(levels, i+1), true)
				return false, prev || curr || next
			}
		}
	}

	return !dampened, dampened
}

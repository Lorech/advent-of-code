package puzzles

import (
	"strings"
)

// Day 25: Code Chronicle
// https://adventofcode.com/2024/day/25
func dayTwentyFive(input string) (int, string) {
	return d25p1(input), ""
}

// Completes the first half of the puzzle for day 25.
func d25p1(input string) int {
	locks, keys := parseDoors(input)
	valid := 0

	for _, lock := range locks {
		for _, key := range keys {
			invalid := false

			for i := 0; i < len(lock); i++ {
				if key[i]+lock[i] >= 6 {
					invalid = true
					break
				}
			}

			if !invalid {
				valid++
			}
		}
	}

	return valid
}

// Parses the input data into structured slices, of lock and key tumbler heights.
func parseDoors(input string) ([][]int, [][]int) {
	rows := strings.Split(input, "\n")
	locks, keys := make([][]int, 0), make([][]int, 0)

	isKey := false
	var obj []int
	for i, row := range rows {
		// Delimiter between locks/keys. Add it to it's slice and reset the object.
		if len(row) == 0 || i == len(rows)-1 {
			if isKey {
				keys = append(keys, obj)
			} else {
				locks = append(locks, obj)
			}
			obj = make([]int, 0)
			continue
		}

		// Start of a new object, so we can determine if it's a lock or key.
		if len(obj) == 0 {
			obj = make([]int, len(row))
			if strings.Contains(row, "#") {
				isKey = false
				continue // The first row doesn't matter for locks.
			} else {
				isKey = true
			}
		}

		// The last row doesn't matter for keys.
		if isKey && len(rows[i+1]) == 0 {
			continue
		}

		// Build up the object column-by-column.
		for j, v := range row {
			if v == '#' {
				obj[j]++
			}
		}
	}

	return locks, keys
}

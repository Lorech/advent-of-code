package aoc2025

import (
	"slices"
	"strings"
)

// Day 7: Laboratories
// https://adventofcode.com/2025/day/7
func daySeven(input string) (int, int) {
	return d7p1(input), 0
}

// Completes the first half of the puzzle for day 7.
func d7p1(input string) int {
	rows, start := strings.Split(input, "\n"), strings.Index(input, "S")
	beams, splits := map[int][]int{0: {start}}, 0

	for y, r := range rows {
		for _, x := range beams[y-1] {
			if r[x] == '^' {
				if x > 0 {
					beams[y] = append(beams[y], x-1)
				}
				if x < len(r) {
					beams[y] = append(beams[y], x+1)
				}
				splits++
			} else {
				beams[y] = append(beams[y], x)
			}
		}

		// Deduplicate to decrease size of next iteration
		slices.Sort(beams[y])
		beams[y] = slices.Compact(beams[y])
	}

	return splits
}

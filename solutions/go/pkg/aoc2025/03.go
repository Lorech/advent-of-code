package aoc2025

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// Day 3: Lobby
// https://adventofcode.com/2025/day/3
func dayThree(input string) (int, int) {
	return d3p1(input), 0
}

// Completes the first half of the puzzle for day 3.
func d3p1(input string) int {
	banks, joltage := strings.Split(input, "\n"), 0

	for _, bank := range banks {
		b := strings.Split(bank, "")
		l := slices.Max(b[:len(b)-1])
		li := slices.Index(b, l)
		r := slices.Max(b[li+1:])
		j, _ := strconv.Atoi(fmt.Sprintf("%s%s", l, r))
		joltage += j
	}

	return joltage
}

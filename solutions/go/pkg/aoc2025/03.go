package aoc2025

import (
	"fmt"
	"lorech/advent-of-code/pkg/convert"
	"slices"
	"strconv"
	"strings"
)

// Day 3: Lobby
// https://adventofcode.com/2025/day/3
func dayThree(input string) (int, int) {
	return d3p1(input), d3p2(input)
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

// Completes the second half of the puzzle for day 3.
func d3p2(input string) int {
	banks, joltage := strings.Split(input, "\n"), 0

	for _, bank := range banks {
		b, j := strings.Split(bank, ""), make([]int, 12)

		for n := 12; n > 0; n-- {
			js := slices.Max(b[:len(b)-n+1])
			i := slices.Index(b, js)
			b = b[i+1:]
			ji, _ := strconv.Atoi(js)
			j[12-n] = ji
		}

		bj, _ := convert.Stoi(j)
		joltage += bj
	}

	return joltage
}

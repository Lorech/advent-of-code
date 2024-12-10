package puzzles

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Day 1: Historian Hysteria
// https://adventofcode.com/2024/day/1
func dayOne(input string) (int, int) {
	return d1p1(input), d1p2(input)
}

// Completes the first half of the puzzle for day 1.
func d1p1(input string) int {
	left, right := parseLists(input)

	distance := 0
	for i := 0; i < len(left); i++ {
		d := right[i] - left[i]
		if d < 0 {
			d = -d
		}
		distance += d
	}

	return distance
}

// Completes the second half of the puzzle for day 1.
func d1p2(input string) int {
	left, right := parseLists(input)

	similarity := 0
	for _, id := range left {
		position, appears := slices.BinarySearch(right, id)
		if appears {
			c := 0
			for i := position; i < len(right); i++ {
				if right[i] != id {
					break
				}

				c++
			}
			similarity += id * c
		}
	}

	return similarity
}

// Parses the two input lists into separate, sorted integer slices.
func parseLists(input string) ([]int, []int) {
	rows := strings.Split(input, "\n")
	left, right := make([]int, len(rows)), make([]int, len(rows))

	for i, row := range rows {
		ids := regexp.MustCompile(`\s+`).Split(row, -1)

		first, _ := strconv.Atoi(ids[0])
		second, _ := strconv.Atoi(ids[1])

		left[i] = first
		right[i] = second
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

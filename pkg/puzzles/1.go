package puzzles

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Day 1: Historian Hysteria
// https://adventofcode.com/2024/day/1
func DayOne(input string) int {
	var left, right []int

	// Parse both lists into separate integer slices.
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		ids := regexp.MustCompile(`\s+`).Split(row, -1)

		if len(ids) != 2 {
			break
		}

		first, error := strconv.Atoi(ids[0])
		if error != nil {
			panic(error)
		}
		second, error := strconv.Atoi(ids[1])
		if error != nil {
			panic(error)
		}

		left = append(left, first)
		right = append(right, second)
	}

	// Sort both lists in ascending order.
	sort.Ints(left)
	sort.Ints(right)

	// Calculate the total distance between the two lists.
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

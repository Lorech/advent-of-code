package puzzles

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Day 1: Historian Hysteria
// https://adventofcode.com/2024/day/1
func DayOne(input string) (int, int) {
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
	slices.Sort(left)
	slices.Sort(right)

	// Calculate the total distance between the two lists.
	distance := 0
	for i := 0; i < len(left); i++ {
		d := right[i] - left[i]
		if d < 0 {
			d = -d
		}
		distance += d
	}

	// Calculate the similarity between the two lists.
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

	return distance, similarity
}

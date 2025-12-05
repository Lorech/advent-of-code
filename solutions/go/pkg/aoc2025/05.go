package aoc2025

import (
	"slices"
	"strconv"
	"strings"
)

type numRange struct {
	min int
	max int
}

// Day 5: Cafeteria
// https://adventofcode.com/2025/day/5
func dayFive(input string) (int, int) {
	return d5p1(input), 0
}

// Completes the first half of the puzzle for day 5.
func d5p1(input string) int {
	ranges, ids := parseInventory(input)
	fresh := 0

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.min && id <= r.max {
				fresh++
				break
			}
		}
	}

	return fresh
}

// Parses the input data into a slice of fresh ranges, and a slice of ingredients.
func parseInventory(input string) ([]numRange, []int) {
	rows := strings.Split(input, "\n")
	ranges, ids := make([]numRange, 0), make([]int, 0)

	splitI := slices.Index(rows, "")

	for i := 0; i < splitI; i++ {
		split := strings.Split(rows[i], "-")
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])
		ranges = append(ranges, numRange{min, max})
	}

	for i := splitI + 1; i < len(rows); i++ {
		id, _ := strconv.Atoi(rows[i])
		ids = append(ids, id)
	}

	return ranges, ids
}

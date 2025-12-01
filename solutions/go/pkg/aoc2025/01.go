package aoc2025

import (
	"strconv"
	"strings"
)

// Day 1: Secret Entrance
// https://adventofcode.com/2025/day/1
func dayOne(input string) (int, int) {
	return d1p1(input), 0
}

// Completes the first half of the puzzle for day 1.
func d1p1(input string) int {
	combination, zeroes := parseLock(input), 0

	position := 50
	for _, move := range combination {
		position = (position + move) % 100

		// In the spirit of the challenge, normalize
		// negative values to underflow from 0 to 99.
		if position < 0 {
			position = 100 + position
		}

		if position == 0 {
			zeroes += 1
		}
	}

	return zeroes
}

// Parses the input data into an integer sequence of moves.
func parseLock(input string) []int {
	rows := strings.Split(input, "\n")
	combination := make([]int, len(rows))

	for i, row := range rows {
		dir, cnt := row[:1], strings.TrimPrefix(row, row[:1])

		num, _ := strconv.Atoi(cnt)

		if dir == "L" {
			num = -num
		}

		combination[i] = num
	}

	return combination
}

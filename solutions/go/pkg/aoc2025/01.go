package aoc2025

import (
	"strconv"
	"strings"
)

// Day 1: Secret Entrance
// https://adventofcode.com/2025/day/1
func dayOne(input string) (int, int) {
	return d1p1(input), d1p2(input)
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

// Completes the second half of the puzzle for day 1.
func d1p2(input string) int {
	combination, position, zeroes := parseLock(input), 50, 0

	for _, move := range combination {
		start := position
		loops := move / 100
		zeroes += max(loops, -loops) // Full loops guarantee passing 0
		move %= 100                  // Exclude the full loops from the move
		position += move

		if position < 0 || position > 99 {
			// Boundary was crossed, so update the values based on the direction
			if position > 99 {
				// Overflow from 99 to 0
				position %= 100
				zeroes += 1
			} else if position < 0 {
				// Underflow from 0 to 99
				position += 100
				// A starting position 0 was counted on the last move
				if start != 0 {
					zeroes += 1
				}
			}
		} else if start != 0 && position == 0 {
			// A boundary was not crossed, so add a zero if we finished the move on it,
			// unless if we started on it, as that was already counted on the last move.
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

package aoc2025

import (
	"fmt"
	"strconv"
	"strings"
)

// Day 2: Gift Shop
// https://adventofcode.com/2025/day/2
func dayTwo(input string) (int, int) {
	return d2p1(input), 0
}

// Completes the first half of the puzzle for day 2.
func d2p1(input string) int {
	ids, sum := parseIds(input), 0

	for _, id := range ids {
		// If the ID consists of an odd amount of digits, it can't consist of two sequences.
		if len(strconv.Itoa(id))%2 == 1 {
			continue
		}

		seq := strconv.Itoa(id % 10)
		for i := id / 10; i > 0; i /= 10 {
			if fmt.Sprintf("%s%s", seq, seq) == strconv.Itoa(id) {
				sum += id
				break
			}

			seq = fmt.Sprintf("%d%s", i%10, seq)
			// If the sequence is longer than the ID, all possibilities are exhausted
			if len(seq) > len(strconv.Itoa(id)) {
				break
			}
		}
	}

	return sum
}

// Parses the input data into an integer sequence of IDs.
func parseIds(input string) []int {
	ranges := strings.Split(input, ",")
	ids := make([]int, 0)

	for _, r := range ranges {
		split := strings.Split(r, "-")

		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])

		for i := start; i <= end; i++ {
			ids = append(ids, i)
		}
	}

	return ids
}

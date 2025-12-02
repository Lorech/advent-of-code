package aoc2025

import (
	"fmt"
	"strconv"
	"strings"
)

// Day 2: Gift Shop
// https://adventofcode.com/2025/day/2
func dayTwo(input string) (int, int) {
	return d2p1(input), d2p2(input)
}

// Completes the first half of the puzzle for day 2.
func d2p1(input string) int {
	ids, sum := parseIds(input), 0

	for _, id := range ids {
		seq := strconv.Itoa(id % 10)
		for i := id / 10; i > 0; i /= 10 {
			if consistsOfSequence(strconv.Itoa(id), seq, 2) {
				sum += id
				break
			}
			seq = fmt.Sprintf("%d%s", i%10, seq)
		}
	}

	return sum
}

// Completes the second half of the puzzle for day 2.
func d2p2(input string) int {
	ids, sum := parseIds(input), 0

	for _, id := range ids {
		seq := strconv.Itoa(id % 10)
		for i := id / 10; i > 0; i /= 10 {
			if consistsOfSequence(strconv.Itoa(id), seq, 50) {
				sum += id
				break
			}
			seq = fmt.Sprintf("%d%s", i%10, seq)
		}
	}

	return sum
}

// Checks if an ID is made up exclusively of sequence seq appearing in the ID
// up to max number of times (inclusive).
func consistsOfSequence(id string, seq string, max int) bool {
	s := seq

	for n := 2; n <= max; n++ {
		s = fmt.Sprintf("%s%s", seq, s)

		// The current sequence size can't fill the ID, but this value
		// flip-flops through iterations, so it may come back around
		if len(id)%2 == 1 && !(len(s)%2 == 1 && n%2 == 1) {
			continue
		}

		// The sequence is too long, so it definitely does not create the ID
		// as any further loops would only make it even longer than it is now
		if len(s) > len(id) {
			return false
		}

		// The sequence matches the ID, which is what we're looking for
		if id == s {
			return true
		}
	}

	return false
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

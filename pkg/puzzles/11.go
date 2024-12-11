package puzzles

import (
	"slices"
	"strconv"
	"strings"
)

// Day 11: Plutonian Pebbles
// https://adventofcode.com/2024/day/11
func dayEleven(input string) (int, int) {
	return d11p1(input), 0
}

// Completes the first half of the puzzle for day 11.
func d11p1(input string) int {
	stones := make([]int, len(strings.Fields(input)))
	for i, s := range strings.Fields(input) {
		stones[i], _ = strconv.Atoi(s)
	}

	for range 25 {
		stones = blink(stones)
	}

	return len(stones)
}

// Handle one blink on the currently visible stones.
func blink(stones []int) []int {
	offset := 0
	for i := range len(stones) {
		j := i + offset
		s := stones[j]
		if s == 0 {
			stones[j] = 1
		} else if n := digits(s); n%2 == 0 {
			half := n / 2
			str := strconv.Itoa(s)
			first, _ := strconv.Atoi(str[:half])
			second, _ := strconv.Atoi(str[half:])
			stones[j] = first
			stones = slices.Insert(stones, j+1, second)
			offset++
		} else {
			stones[j] *= 2024
		}
	}

	return stones
}

// Get the number of digits from an integer.
//
// NOTE: A good function to extract into a utility.
func digits(i int) int {
	if i == 0 {
		return 1
	}

	n := 0
	for i != 0 {
		i /= 10
		n++
	}

	return n
}

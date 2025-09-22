package puzzles

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

// Day 22: Monkey Market
// https://adventofcode.com/2024/day/22
func dayTwentyTwo(input string) (int, int) {
	return d22p1(input), d22p2(input)
}

// Completes the first half of the puzzle for day 22.
func d22p1(input string) int {
	secrets := parsePrices(input)
	sum := 0

	for _, secret := range secrets {
		for range 2000 {
			secret = newSecret(secret)
		}
		sum += secret
	}

	return sum
}

// Completes the second half of the puzzle for day 22.
func d22p2(input string) int {
	secrets := parsePrices(input)
	bananas := 0
	sequences := make(map[[4]int]int)

	for _, secret := range secrets {
		deltas := []int{0}
		secretSequences := make([][4]int, 0)
		for range 2000 {
			p := secret % 10
			secret = newSecret(secret)
			n := secret % 10
			deltas = append(deltas, n-p)

			if len(deltas) == 4 {
				sequence := [4]int{deltas[0], deltas[1], deltas[2], deltas[3]}
				// Only count the first appearance of a sequence.
				if !slices.Contains(secretSequences, sequence) {
					sequences[sequence] += n
					secretSequences = append(secretSequences, sequence)
				}
				deltas = deltas[1:]
			}
		}
	}

	for _, total := range sequences {
		if total > bananas {
			bananas = total
		}
	}

	return bananas
}

// Derives a new secret number based on an existing secret number.
func newSecret(o int) int {
	n := o

	n = ((n * 64) ^ n) % 16777216
	n = ((int(math.Floor(float64(n) / float64(32)))) ^ n) % 16777216
	n = ((n * 2048) ^ n) % 16777216

	return n
}

// Parses the input data into the initial integer values of the secret numbers.
func parsePrices(input string) []int {
	rows := strings.Split(input, "\n")
	nums := make([]int, len(rows))
	for i, row := range rows {
		num, _ := strconv.Atoi(row)
		nums[i] = num
	}
	return nums
}

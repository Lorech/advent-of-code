package puzzles

import (
	"slices"
	"strconv"
	"strings"
)

type stone struct {
	count          int   // The number of stones with this value.
	transformation []int // The values of the resulting transformation of this stone after a blink.
}

// Day 11: Plutonian Pebbles
// https://adventofcode.com/2024/day/11
func dayEleven(input string) (int, int) {
	return d11p1(input), d11p2(input)
}

// Completes the first half of the puzzle for day 11.
// TODO: Refactor this to use the new apprach. Leaving it until I benchmark it for comparsison.
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

// Completes the second half of the puzzle for day 11.
func d11p2(input string) int {
	stones := parseStones(input)       // Metadata and cache for stones.
	values := make([]int, len(stones)) // Unique stone values for quicker lookup.
	i := 0
	for k := range stones {
		values[i] = k
		i++
	}

	for range 75 {
		updates := make(map[int]int, 0) // Track the delta between the previous and current blink.

		for _, val := range values {
			o, _ := stones[val]
			updates[val] -= o.count

			// Update each of the known transformations.
			for _, n := range o.transformation {
				updates[n] += o.count
			}
		}

		// Update the stone data post-blink.
		values = make([]int, 0)
		for val, delta := range updates {
			s, exists := stones[val]

			if exists {
				s.count += delta
				stones[val] = s
			} else {
				stones[val] = stone{delta, transform(val)}
			}

			if stones[val].count > 0 {
				values = append(values, val)
			}
		}
	}

	n := 0
	for _, val := range values {
		s, _ := stones[val]
		n += s.count
	}

	return n
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

// Transform the value of a stone into its new representation after a blink.
func transform(i int) []int {
	if i == 0 {
		return []int{1}
	} else if n := digits(i); n%2 == 0 {
		half := n / 2
		str := strconv.Itoa(i)
		first, _ := strconv.Atoi(str[:half])
		second, _ := strconv.Atoi(str[half:])
		return []int{first, second}
	} else {
		return []int{i * 2024}
	}
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

// Parses the input data to generate the initial batch of stone values.
func parseStones(input string) map[int]stone {
	stones := make(map[int]stone, len(strings.Fields(input)))

	for _, v := range strings.Fields(input) {
		val, _ := strconv.Atoi(v)
		s, exists := stones[val]

		if exists {
			s.count++
		} else {
			stones[val] = stone{1, transform(val)}
		}
	}

	return stones
}

package aoc2025

import (
	"regexp"
	"strings"
)

// Day 11: Reactor
// https://adventofcode.com/2025/day/11
func dayEleven(input string) (int, int) {
	return d11p1(input), d11p2(input)
}

// Completes the first half of the puzzle for day 11.
func d11p1(input string) int {
	racks := parseRacks(input)
	paths := 0

	stack := racks["you"]
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if s == "out" {
			paths++
			continue
		}

		for _, o := range racks[s] {
			stack = append(stack, o)
		}
	}

	return paths
}

// Completes the second half of the puzzle for day 11.
func d11p2(input string) int {
	racks := parseRacks(input)

	// Option A: SVR -> DAC -> FFT -> OUT
	a1, a2, a3 := 0, 0, 0
	a1c, a2c, a3c := map[string]int{}, map[string]int{}, map[string]int{}
	a1 = traverseRacks("svr", "dac", racks, &a1c)
	a2 = traverseRacks("dac", "fft", racks, &a2c)
	a3 = traverseRacks("fft", "out", racks, &a3c)
	a := a1 * a2 * a3

	// Option B: SVR -> FFT -> DAC -> OUT
	b1, b2, b3 := 0, 0, 0
	b1c, b2c, b3c := map[string]int{}, map[string]int{}, map[string]int{}
	b1 = traverseRacks("svr", "fft", racks, &b1c)
	b2 = traverseRacks("fft", "dac", racks, &b2c)
	b3 = traverseRacks("dac", "out", racks, &b3c)
	b := b1 * b2 * b3

	return int(a + b)
}

// Navigates the graph, returning the number of steps between `start` and `end`.
// A local cache is used to avoid re-proccesing already computed nodes.
// TODO: Integrate this into part 1 as performance optimization.
func traverseRacks(start, end string, racks map[string][]string, cache *map[string]int) int {
	if start == end {
		return 1
	}

	steps, cached := (*cache)[start]
	if cached {
		return steps
	}

	result := 0
	for _, c := range racks[start] {
		result += traverseRacks(c, end, racks, cache)
	}

	(*cache)[start] = result
	return result
}

// Parses input data into structured data.
func parseRacks(input string) map[string][]string {
	rows, racks := strings.Split(input, "\n"), map[string][]string{}
	re := regexp.MustCompile(`\w+`)

	for _, r := range rows {
		ps := re.FindAllString(r, -1)
		racks[ps[0]] = make([]string, len(ps)-1)
		for i, o := range ps[1:] {
			racks[ps[0]][i] = o
		}
	}

	return racks
}

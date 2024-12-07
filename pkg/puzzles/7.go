package puzzles

import (
	"strconv"
	"strings"
)

// Day 7: Bridge Repair
// https://adventofcode.com/2024/day/7
func daySeven(input string) (int, int) {
	return d7p1(input), 0
}

// Completes the first half of the puzzle for day 7.
func d7p1(input string) int {
	equations := parse(input)
	sum := 0

	for result, components := range equations {
		if produce(components[0], result, components[1:]) {
			sum += result
		}
	}

	return sum
}

// Recursively constructs every possible equation from the provided components
// and determines if it is possible to insert operands in a way to obtain the result.
func produce(current int, expected int, components []int) bool {
	// Every possibility in this tree has been processed, did we get the right result?
	if len(components) == 0 {
		return current == expected
	}

	// We overshot - no point in continuing.
	if current > expected {
		return false
	}

	sum := produce(current+components[0], expected, components[1:])
	mul := produce(current*components[0], expected, components[1:])

	return sum || mul
}

// Parses the input data into a map which matches the structure of the input:
// - The key is the expected result of the equation;
// - The value is a slice containing all number parts of the equation.
func parse(input string) map[int][]int {
	mapping := make(map[int][]int, 0)
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		// Crappy file parsing strikes again.
		if row == "" {
			continue
		}

		components := strings.Split(row, ": ")
		result, _ := strconv.Atoi(components[0])
		parts := strings.Split(components[1], " ")
		mapping[result] = make([]int, len(parts))
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			mapping[result][i] = num
		}
	}

	return mapping
}

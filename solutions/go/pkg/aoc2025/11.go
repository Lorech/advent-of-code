package aoc2025

import (
	"regexp"
	"strings"
)

// Day 11: Reactor
// https://adventofcode.com/2025/day/10
func dayEleven(input string) (int, int) {
	return d11p1(input), 0
}

// Completes the first half of the puzzle for day 11.
func d11p1(input string) int {
	racks := parseRacks(input)
	paths := 0

	stack, visited := racks["you"], []string{}
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited = append(visited, s)

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

package aoc2025

import (
	"math"
	"slices"
	"strings"
)

// Day 7: Laboratories
// https://adventofcode.com/2025/day/7
func daySeven(input string) (int, int) {
	return d7p1(input), d7p2(input)
}

// Completes the first half of the puzzle for day 7.
func d7p1(input string) int {
	rows, start := strings.Split(input, "\n"), strings.Index(input, "S")
	beams, splits := map[int][]int{0: {start}}, 0

	for y, r := range rows {
		for _, x := range beams[y-1] {
			ts := make([]int, 0)
			if r[x] == '^' {
				if x > 0 {
					ts = append(ts, x-1)
				}
				if x < len(r) {
					ts = append(ts, x+1)
				}
				splits++
			} else {
				ts = append(ts, x)
			}

			for _, xn := range ts {
				repeat := slices.Contains(beams[y], xn)
				if !repeat {
					beams[y] = append(beams[y], xn)
				}
			}
		}
	}

	return splits
}

// Completes the second half of the puzzle for day 7.
func d7p2(input string) int {
	rows, start := strings.Split(input, "\n"), strings.Index(input, "S")
	beams, apps := map[int][]int{0: {start}}, map[int]map[int]int{}
	for y := range rows {
		apps[y] = map[int]int{}
	}

	for y, r := range rows {
		for _, x := range beams[y-1] {
			ts := make([]int, 0)
			if r[x] == '^' {
				if x > 0 {
					ts = append(ts, x-1)
				}
				if x < len(r) {
					ts = append(ts, x+1)
				}
			} else {
				ts = append(ts, x)
			}

			for _, xn := range ts {
				_, repeat := apps[y][xn]
				if repeat {
					apps[y][xn] += int(math.Max(1, float64(apps[y-1][x])))
				} else {
					apps[y][xn] = int(math.Max(1, float64(apps[y-1][x])))
					beams[y] = append(beams[y], xn)
				}
			}
		}
	}

	count := 0
	for _, n := range apps[len(rows)-1] {
		count += n
	}
	return count
}

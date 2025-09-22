package puzzles

import (
	"slices"
	"strconv"
	"strings"
)

// Day 10: Hoof It
// https://adventofcode.com/2024/day/10
func dayTen(input string) (int, int) {
	return d10p1(input), d10p2(input)
}

// Completes the first half of the puzzle for day 10.
func d10p1(input string) int {
	trail, heads := parseTrail(input)
	peaks := make([][2]int, 0)

	for _, head := range heads {
		p := walkTrail(head, trail)

		// Filter only distinct peaks that can be reached from this head.
		for i, peak := range p {
			if slices.Contains(p[i+1:], peak) {
				continue
			}

			peaks = append(peaks, peak)
		}
	}

	return len(peaks)
}

// Completes the second half of the puzzle for day 10.
func d10p2(input string) int {
	trail, heads := parseTrail(input)
	peaks := make([][2]int, 0)

	for _, head := range heads {
		peaks = append(peaks, walkTrail(head, trail)...)
	}

	return len(peaks)
}

// Recursively walks the full trail from the starting position found in `tile`.
// Returns a slice of all peaks that can be reached from this position.
func walkTrail(tile [2]int, trail [][]int) [][2]int {
	x, y := tile[0], tile[1]

	// This is the end of the trail!
	if trail[y][x] == 9 {
		return [][2]int{{x, y}}
	}

	peaks := make([][2]int, 0)

	if y-1 >= 0 && trail[y-1][x] == trail[y][x]+1 {
		peaks = append(peaks, walkTrail([2]int{x, y - 1}, trail)...)
	}

	if y+1 < len(trail) && trail[y+1][x] == trail[y][x]+1 {
		peaks = append(peaks, walkTrail([2]int{x, y + 1}, trail)...)
	}

	if x-1 >= 0 && trail[y][x-1] == trail[y][x]+1 {
		peaks = append(peaks, walkTrail([2]int{x - 1, y}, trail)...)
	}

	if x+1 < len(trail[0]) && trail[y][x+1] == trail[y][x]+1 {
		peaks = append(peaks, walkTrail([2]int{x + 1, y}, trail)...)
	}

	return peaks
}

// Parses the input data, returning the full trail map split at y-coordinates,
// and the x,y coordinates of every trailhead on the map.
func parseTrail(input string) ([][]int, [][2]int) {
	rows := strings.Split(input, "\n")
	trail := make([][]int, len(rows))
	heads := make([][2]int, 0)

	for y, row := range rows {
		trail[y] = make([]int, len(row))
		for x, tile := range row {
			trail[y][x], _ = strconv.Atoi(string(tile))
			if tile == '0' {
				heads = append(heads, [2]int{x, y})
			}
		}
	}

	return trail, heads
}

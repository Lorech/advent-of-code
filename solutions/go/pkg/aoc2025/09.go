package aoc2025

import (
	"lorech/advent-of-code/pkg/grid"
	"strconv"
	"strings"
)

// Day 9: Movie Theater
// https://adventofcode.com/2025/day/9
func dayNine(input string) (int, int) {
	return d9p1(input), 0
}

// Completes the first half of the puzzle for day 9.
func d9p1(input string) int {
	tiles, maxArea := parseTiles(input), 0

	for i := 0; i < len(tiles)-1; i++ {
		a := tiles[i]
		for j := i + 1; j < len(tiles); j++ {
			b := tiles[j]
			x, y := a.X-b.X, a.Y-b.Y
			if x < 0 {
				x *= -1
			}
			if y < 0 {
				y *= -1
			}
			// Off-by-one due to subtraction
			x++
			y++
			area := x * y
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

// Parses input data into structured data.
func parseTiles(input string) []grid.Coordinate {
	rows := strings.Split(input, "\n")
	tiles := make([]grid.Coordinate, len(rows))
	for i, r := range rows {
		c := strings.Split(r, ",")
		x, _ := strconv.Atoi(c[0])
		y, _ := strconv.Atoi(c[1])
		tiles[i] = grid.Coordinate{X: x, Y: y}
	}
	return tiles
}

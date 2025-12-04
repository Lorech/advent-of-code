package aoc2025

import (
	"fmt"
	"strings"
)

// Day 4: Printing Department
// https://adventofcode.com/2025/day/4
func dayFour(input string) (int, int) {
	return d4p1(input), d4p2(input)
}

// Completes the first half of the puzzle for day 4.
func d4p1(input string) int {
	grid, accessible := strings.Split(input, "\n"), 0

	for y, row := range grid {
		for x, tile := range row {
			if tile != '@' {
				continue
			}

			adjacent := 0
			for yn := -1; yn <= 1; yn++ {
				for xn := -1; xn <= 1; xn++ {
					if yn == 0 && xn == 0 {
						continue
					}

					if y+yn < 0 || y+yn >= len(grid) || x+xn < 0 || x+xn >= len(row) {
						continue
					}

					if grid[y+yn][x+xn] == '@' {
						adjacent += 1
					}

					if adjacent == 4 {
						break
					}
				}

				if adjacent == 4 {
					break
				}
			}

			if adjacent < 4 {
				accessible += 1
			}
		}
	}

	return accessible
}

// Completes the second half of the puzzle for day 4.
func d4p2(input string) int {
	grid, accessible, removed := strings.Split(input, "\n"), 0, make([][]int, 0)

	for accessible == 0 || len(removed) > 0 {
		for _, tile := range removed {
			grid[tile[0]] = fmt.Sprintf("%s%s%s", grid[tile[0]][:tile[1]], ".", grid[tile[0]][tile[1]+1:])
		}
		removed = [][]int{}

		for y, row := range grid {
			for x, tile := range row {
				if tile != '@' {
					continue
				}

				adjacent := 0
				for yn := -1; yn <= 1; yn++ {
					for xn := -1; xn <= 1; xn++ {
						if yn == 0 && xn == 0 {
							continue
						}

						if y+yn < 0 || y+yn >= len(grid) || x+xn < 0 || x+xn >= len(row) {
							continue
						}

						if grid[y+yn][x+xn] == '@' {
							adjacent += 1
						}

						if adjacent == 4 {
							break
						}
					}

					if adjacent == 4 {
						break
					}
				}

				if adjacent < 4 {
					accessible += 1
					removed = append(removed, []int{y, x})
				}
			}
		}
	}

	return accessible
}

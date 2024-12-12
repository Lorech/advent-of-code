package puzzles

import "strings"

type plant struct {
	X         int
	Y         int
	Kind      rune
	Neighbors int
}

// Day 12: Garden Groups
// https://adventofcode.com/2024/day/12
func dayTwelve(input string) (int, int) {
	return d12p1(input), 0
}

// Completes the first half of the puzzle for day 12.
func d12p1(input string) int {
	garden := parseGarden(input)
	plots := make([][]*plant, 0)
	visited := make(map[[2]int]bool)
	price := 0

	// Collect all unique plots of plants.
	for y, row := range garden {
		for x, tile := range row {
			_, v := visited[[2]int{x, y}]
			if !v {
				start := plant{x, y, tile, 0}
				plot := make([]*plant, 0)
				fencePlot(&start, &plot, &visited, garden)
				plots = append(plots, plot)
			}
		}
	}

	// Calculate the price of fencing the plot.
	for _, plot := range plots {
		area := len(plot)
		perimiter := 0
		for _, p := range plot {
			perimiter += 4 - p.Neighbors
		}
		price += area * perimiter
	}

	return price
}

// Recursively navigate the plot of land matching the starting plant.
func fencePlot(start *plant, plot *[]*plant, visited *map[[2]int]bool, garden [][]rune) {
	x, y := start.X, start.Y
	value := garden[y][x]
	*plot = append(*plot, start)
	(*visited)[[2]int{x, y}] = true

	directions := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, direction := range directions {
		nextX, nextY := x+direction[0], y+direction[1]

		// Out of bounds.
		if nextX < 0 || nextX >= len((garden)[0]) || nextY < 0 || nextY >= len(garden) {
			continue
		}

		// Add a neighbor if it's the same plant or skip it if not.
		if garden[nextY][nextX] == value {
			start.Neighbors++
		} else {
			continue
		}

		// Skip if we've been on this field.
		if (*visited)[[2]int{nextX, nextY}] {
			continue
		}

		// Check the new plant.
		next := &plant{nextX, nextY, garden[nextY][nextX], 0}
		fencePlot(next, plot, visited, garden)
	}
}

// Parses the input data into a grid of runes.
func parseGarden(input string) [][]rune {
	rows := strings.Split(input, "\n")
	garden := make([][]rune, len(rows))
	for i, row := range rows {
		garden[i] = []rune(row)
	}
	return garden
}

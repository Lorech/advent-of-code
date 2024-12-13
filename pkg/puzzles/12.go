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
	return d12p1(input), d12p2(input)
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
				fencePlot(&start, &plot, nil, &visited, garden)
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

// Completes the second half of the puzzle for day 12.
func d12p2(input string) int {
	garden := parseGarden(input)
	visited := make(map[[2]int]bool)
	price := 0

	for y, row := range garden {
		for x, tile := range row {
			_, v := visited[[2]int{x, y}]
			if !v {
				start := plant{x, y, tile, 0}
				plot := make([]*plant, 0)
				edges := 0
				fencePlot(&start, &plot, &edges, &visited, garden)
				area := len(plot)
				price += area * edges
			}
		}
	}

	return price
}

// Recursively navigate the plot of land matching the starting plant.
// NOTE: Ugly function just to get a result. Needs to be cleaned up and split
// between parts a bit better if I ever return to optimize.
func fencePlot(start *plant, plot *[]*plant, edges *int, visited *map[[2]int]bool, garden [][]rune) {
	x, y := start.X, start.Y
	value := garden[y][x]
	*plot = append(*plot, start)
	(*visited)[[2]int{x, y}] = true

	// If we have the pointer to edges, we should check the corners!
	if edges != nil {
		checkCorners(start, edges, garden)
	}

	// Flood-fill the plot while also calculating neighbors for each tile to get the perimeter.
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
		fencePlot(next, plot, edges, visited, garden)
	}
}

// Checks the corners of the provided tile to increment the number of edges
// that the plant plot has.
func checkCorners(start *plant, edges *int, garden [][]rune) {
	x, y := start.X, start.Y
	value := garden[y][x]
	gW, gH := len(garden[0]), len(garden)

	corners := [4][3][2]int{
		{{0, -1}, {1, -1}, {1, 0}},   // Top-right
		{{1, 0}, {1, 1}, {0, 1}},     // Bottom-right
		{{0, 1}, {-1, 1}, {-1, 0}},   // Bottom-left
		{{-1, 0}, {-1, -1}, {0, -1}}, // Top-left
	}
	for _, corner := range corners {
		aX, aY := x+corner[0][0], y+corner[0][1]
		bX, bY := x+corner[1][0], y+corner[1][1]
		cX, cY := x+corner[2][0], y+corner[2][1]

		// Outside corner if neighbors are out of bounds or a different plant.
		if aX < 0 || aX >= gW || aY < 0 || aY >= gH || garden[aY][aX] != value {
			if cX < 0 || cX >= gW || cY < 0 || cY >= gH || garden[cY][cX] != value {
				// This is an outside corner! That also means it can't be an inside corner.
				*edges++
				continue
			}
		}

		// Inside corner if neighbors are the same plant, but the diagonal is different.
		if aX >= 0 && aX < gW && aY >= 0 && aY < gH && garden[aY][aX] == value {
			if bX >= 0 && bX < gW && bY >= 0 && bY < gH && garden[bY][bX] != value {
				if cX >= 0 && cX < gW && cY >= 0 && cY < gH && garden[cY][cX] == value {
					// This is an inside corner!
					*edges++
				}
			}
		}
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

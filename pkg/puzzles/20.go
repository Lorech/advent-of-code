package puzzles

import (
	"lorech/advent-of-code-2024/pkg/grid"
	"strings"
)

// Day 20: Race Condition
// https://adventofcode.com/2024/day/20
func dayTwenty(input string) (interface{}, interface{}) {
	return d20p1(input), 0
}

// Completes the second half of the puzzle for day 20.
func d20p1(input string, options ...int) int {
	threshold := 100
	if len(options) > 0 {
		threshold = options[0]
	}

	maze, start, end := parseMaze(input)
	e, _ := grid.NavigateMaze(maze, start, end)
	path := make([]grid.Tile, 0)
	for e.Parent != nil {
		path = append(path, e)
		e = *e.Parent
	}
	path = append(path, e) // Add the start node to the path.

	savings := make(map[int]int, 0)
	for i := 0; i < len(path)-4; i++ {
		t1 := path[i]
		for j := i + 4; j < len(path); j++ {
			t2 := path[j]
			dy, dx := t2.Position.Y-t1.Position.Y, t2.Position.X-t1.Position.X
			dy = max(dy, -dy)
			dx = max(dx, -dx)
			if dy == 2 && dx == 0 || dy == 0 && dx == 2 {
				time := j - i - 2
				savings[time]++
			}
		}
	}

	result := 0
	for savings, count := range savings {
		if savings >= threshold {
			result += count
		}
	}
	return result
}

// Parses the input data into structured data:
//   - The maze itself;
//   - The coordinates of the original starting position;
//   - The coordinates of the original ending position.
func parseMaze(input string) ([][]rune, grid.Coordinates, grid.Coordinates) {
	rows := strings.Split(input, "\n")
	maze := make([][]rune, len(rows))
	var start, end grid.Coordinates

	for y, row := range rows {
		maze[y] = make([]rune, len(row))
		for x, tile := range row {
			switch tile {
			case 'S':
				start = grid.Coordinates{X: x, Y: y}
				maze[y][x] = '.'
			case 'E':
				end = grid.Coordinates{X: x, Y: y}
				maze[y][x] = '.'
			case '#':
				maze[y][x] = '#'
			case '.':
				maze[y][x] = '.'
			}
		}
	}

	return maze, start, end
}

package puzzles

import (
	"lorech/advent-of-code-2024/pkg/grid"
	"strings"
)

// Day 20: Race Condition
// https://adventofcode.com/2024/day/20
func dayTwenty(input string) (int, int) {
	return d20p1(input), d20p2(input)
}

// Completes the first half of the puzzle for day 20.
// TODO: Benchmark and compare to part 2. Part 2 has a reusable solution, which
// may also be more efficient and speed up part 1 if applied.
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

// Completes the second half of the puzzle for day 20.
func d20p2(input string, options ...int) int {
	threshold := 100
	if len(options) > 0 {
		threshold = options[0]
	}

	maze, start, end := parseMaze(input)
	e, _ := grid.NavigateMaze(maze, start, end)
	path := make([][2]int, 0)
	for e.Parent != nil {
		path = append([][2]int{{e.Position.Y, e.Position.X}}, path...)
		e = *e.Parent
	}
	path = append([][2]int{{e.Position.Y, e.Position.X}}, path...) // Add the start node to the path.

	savings := 0
	for i, a := range path[:len(path)-threshold] {
		for j, b := range path[i+threshold:] {
			distance := grid.ManhattanDistance(
				grid.Coordinates{X: a[1], Y: a[0]},
				grid.Coordinates{X: b[1], Y: b[0]},
			)
			if distance <= 20 && distance <= j {
				savings++
			}
		}
	}

	return savings
}

// Find the unique tiles that can be shortcut from the current tile in the
// provided time. This does not necessarily mean that this shortcut will save
// time on a full run - just that it is reachable.
func findShortcuts(tile [2]int, maze [][]rune, time int) [][2]int {
	exits := make([][2]int, 0)
	options := grid.WithinManhattanDistance(grid.Coordinates{X: tile[1], Y: tile[0]}, time)
	for _, option := range options {
		if option.Y < 0 || option.Y >= len(maze) || option.X < 0 || option.X >= len(maze[0]) {
			continue
		}

		if maze[option.Y][option.X] == '.' {
			exits = append(exits, [2]int{option.Y, option.X})
		}
	}
	return exits
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

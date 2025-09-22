package puzzles

import (
	"lorech/advent-of-code/pkg/grid"
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

	maze, start, end := parseTrack(input)
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

	maze, start, end := parseTrack(input)
	e, _ := grid.NavigateMaze(maze, start, end)
	path := make([]grid.Coordinates, 0)
	for e.Parent != nil {
		path = append([]grid.Coordinates{e.Position}, path...)
		e = *e.Parent
	}
	path = append([]grid.Coordinates{e.Position}, path...) // Add the start node to the path.

	savings := 0
	for i, a := range path[:len(path)-threshold] {
		for j, b := range path[i+threshold:] { // j - time saved above threshold by teleporting.
			distance := grid.ManhattanDistance(a, b)
			// The tile is close enough that we don't exceed cheat time, and we also
			// remain over the threshold of time save that we want to enumerate.
			if distance <= 20 && distance <= j {
				savings++
			}
		}
	}

	return savings
}

// Parses the input data into structured data:
//   - The maze itself;
//   - The coordinates of the original starting position;
//   - The coordinates of the original ending position.
func parseTrack(input string) ([][]rune, grid.Coordinates, grid.Coordinates) {
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

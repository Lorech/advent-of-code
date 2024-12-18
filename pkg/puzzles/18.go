package puzzles

import (
	"lorech/advent-of-code-2024/pkg/grid"
	"slices"
	"strconv"
	"strings"
)

type node struct {
	position grid.Coordinates // The position of this node.
	parent   *node            // The parent node for obtaining the shortest path to this node.
}

// Day 18: RAM Run
// https://adventofcode.com/2024/day/18
func dayEighteen(input string) (int, int) {
	return d18p1(input), 0
}

// Completes the first half of the puzzle for day 18.
// As test data uses different inputs, a slice of optional integer parameters
// can be specified with up to 3 values:
//   - width of the maze
//   - height of the maze
//   - number of blocks dropped onto the maze
//
// The default of these variables match the puzzle requirements (71, 71, 1024).
func d18p1(input string, options ...int) int {
	w, h, n := 71, 71, 1024
	if len(options) > 0 {
		w = options[0]
		h = options[1]
		n = options[2]
	}

	maze := createMaze(input, w, h, n)
	end, success := navigateMaze(maze)
	if !success {
		panic("Failed to reach the end.")
	}

	path := make([]grid.Coordinates, 0)
	for end.parent != nil {
		path = append(path, end.position)
		end = *end.parent
	}

	return len(path)
}

// Navigates the maze using BFS, returning the finishing node, whose parent
// nodes, chained all the way to the top, will form the shortest path to reach
// the end. A second return value is provided as an indicator if the end was
// actually reached.
func navigateMaze(maze [][]rune) (node, bool) {
	q := make([]node, 0)
	v := make([][]bool, len(maze))
	for i := range v {
		v[i] = make([]bool, len(maze[0]))
	}

	q = append(q, node{grid.Coordinates{X: 0, Y: 0}, nil})
	v[0][0] = true

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if n.position.Y == len(maze)-1 && n.position.X == len(maze[0])-1 {
			return n, true
		}

		dirs := [4]grid.Direction{grid.Up, grid.Down, grid.Left, grid.Right}
		for _, dir := range dirs {
			yd, xd := dir.Velocity()
			ny, nx := n.position.Y+yd, n.position.X+xd
			if ny >= 0 && ny < len(maze) && nx >= 0 && nx < len(maze[0]) && !v[ny][nx] && maze[ny][nx] != '#' {
				v[ny][nx] = true
				q = append(q, node{grid.Coordinates{X: nx, Y: ny}, &n})
			}
		}
	}

	return node{}, false
}

// Prepares the maze based on the provided input data, returning a 2D slice of
// runes, in the specified height and width of the maze, laying the amount of
// obstacles as specified in the obstacles parameter.
func createMaze(input string, width int, height int, obstacles int) [][]rune {
	maze := make([][]rune, height)
	for y := range height {
		maze[y] = slices.Repeat([]rune{'.'}, width)
	}

	rows := strings.Split(input, "\n")
	for i := 0; i < obstacles; i++ {
		coords := strings.Split(rows[i], ",")
		ox, _ := strconv.Atoi(coords[0])
		oy, _ := strconv.Atoi(coords[1])
		maze[oy][ox] = '#'
	}

	return maze
}

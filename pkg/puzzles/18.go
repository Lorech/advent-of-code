package puzzles

import (
	"fmt"
	"lorech/advent-of-code-2024/pkg/grid"
	"math"
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
func dayEighteen(input string) (interface{}, interface{}) {
	return d18p1(input), d18p2(input)
}

// Completes the first half of the puzzle for day 18.
//
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

	obstacles := parseObstacles(input)
	maze := createMaze(w, h, obstacles[:n])
	end, _ := navigateMaze(maze)

	path := make([]grid.Coordinates, 0)
	for end.parent != nil {
		path = append(path, end.position)
		end = *end.parent
	}

	return len(path)
}

// Completes the second half of the puzzle for day 18.
//
// As test data uses different inputs, a slice of optional integer parameters
// can be specified with up to 2 values:
//   - width of the maze
//   - height of the maze
//
// The default of these variables match the puzzle requirements (71, 71).
func d18p2(input string, options ...int) string {
	w, h := 71, 71
	if len(options) > 0 {
		w = options[0]
		h = options[1]
	}

	obstacles := parseObstacles(input)
	left, right := 0, len(obstacles)-1
	middle := int(math.Floor(float64((left + right) / 2)))
	bounds := []int{left, middle, right}
	var result string

	for true {
		maze := createMaze(w, h, obstacles[left:middle])
		_, success := navigateMaze(maze)
		mStartIndex, _ := slices.BinarySearch(bounds, middle)
		mEndIndex := mStartIndex

		// Find the middle value for the next iteration based on binary search.
		if success {
			mEndIndex += 1

		} else {
			mEndIndex -= 1
		}
		middle = int(math.Floor(float64((middle + bounds[mEndIndex]) / 2)))

		// We've already seen this value before, which means we know the bound.
		if slices.Contains(bounds, middle) {
			maze := createMaze(w, h, obstacles[left:middle])
			_, success := navigateMaze(maze)
			var byte [2]int
			if success {
				byte = obstacles[middle]
			} else {
				byte = obstacles[middle+1]
			}
			result = fmt.Sprintf("%d,%d", byte[1], byte[0])
			break
		}

		// Prepare the bounds for the next loop and go again.
		bounds = append(bounds, middle)
		slices.Sort(bounds)
	}

	return result
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

// Prepares a maze based on the provided size, and populates it with the
// provided obstacles. The function assumes that the obstacles are within the
// bounds of the provided maze size, crashing on an index error if one occurs.
func createMaze(width int, height int, obstacles [][2]int) [][]rune {
	maze := make([][]rune, height)
	for y := range height {
		maze[y] = slices.Repeat([]rune{'.'}, width)
	}

	for _, obstacle := range obstacles {
		maze[obstacle[0]][obstacle[1]] = '#'
	}

	return maze
}

// Parses the provided input data into a slice of y,x coordinates where
// obstacles may be found within the maze.
func parseObstacles(input string) [][2]int {
	rows := strings.Split(input, "\n")
	obstacles := make([][2]int, len(rows))
	for i, row := range rows {
		coords := strings.Split(row, ",")
		ox, _ := strconv.Atoi(coords[0])
		oy, _ := strconv.Atoi(coords[1])
		obstacles[i] = [2]int{oy, ox}
	}
	return obstacles
}

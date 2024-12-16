package puzzles

import (
	"lorech/advent-of-code-2024/pkg/grid"
	"math"
	"strings"
)

// Day 16: Reindeer Maze
// https://adventofcode.com/2024/day/16
func daySixteen(input string) (int, int) {
	return d16p1(input), 0
}

// Completes the first half of the puzzle for day 16.
func d16p1(input string) int {
	maze, start := parseMaze(input)
	visited := make([][]bool, len(maze))
	for y := range visited {
		visited[y] = make([]bool, len(maze[0]))
	}
	score := 0
	minScore := math.MaxInt
	navigateMaze(start, grid.Right, maze, &visited, &score, &minScore)
	return minScore
}

// Traverses the maze, looking for the path that produces the lowest final score.
func navigateMaze(position [2]int, direction grid.Direction, maze [][]rune, visited *[][]bool, score *int, minScore *int) {
	y, x := position[0], position[1]

	// Skip this tile because we've been here or it's a wall.
	if maze[y][x] == '#' || (*visited)[y][x] {
		return
	}

	// This path can't be better than one we've found already.
	if *score >= *minScore {
		return
	}

	// This is the end, so we may have a new best path.
	if maze[y][x] == 'E' {
		if *score < *minScore {
			*minScore = *score
		}
		return
	}

	(*visited)[y][x] = true

	// Prefer navigating in the direction of travel to find the cheaper path.
	dirs := [4]grid.Direction{grid.Up, grid.Left, grid.Right, grid.Down}
	switch direction {
	case grid.Down:
		dirs = [4]grid.Direction{grid.Down, grid.Left, grid.Right, grid.Up}
	case grid.Left:
		dirs = [4]grid.Direction{grid.Left, grid.Up, grid.Down, grid.Right}
	case grid.Right:
		dirs = [4]grid.Direction{grid.Right, grid.Up, grid.Down, grid.Left}
	}

	for _, dir := range dirs {
		yd, xd := dir.Velocity()
		moveScore := 1

		// Change the move score if we have to turn.
		switch direction {
		case grid.Up:
			if dir == grid.Left || dir == grid.Right {
				moveScore += 1000
			} else if dir == grid.Down {
				moveScore += 2000
			}
		case grid.Down:
			if dir == grid.Left || dir == grid.Right {
				moveScore += 1000
			} else if dir == grid.Up {
				moveScore += 2000
			}
		case grid.Left:
			if dir == grid.Up || dir == grid.Down {
				moveScore += 1000
			} else if dir == grid.Right {
				moveScore += 2000
			}
		case grid.Right:
			if dir == grid.Up || dir == grid.Down {
				moveScore += 1000
			} else if dir == grid.Left {
				moveScore += 2000
			}
		}

		*score += moveScore
		navigateMaze([2]int{y + yd, x + xd}, dir, maze, visited, score, minScore)
		*score -= moveScore
	}

	// Backtrack cause we've done all this tile can do.
	(*visited)[y][x] = false
}

// Parses the input data into a 2D slice of runes representing the maze, and
// returns it along with the starting position of the reindeer in the maze.
func parseMaze(input string) ([][]rune, [2]int) {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows))
	var reindeer [2]int

	for y, row := range rows {
		grid[y] = make([]rune, len(row))
		grid[y] = []rune(row)

		if reindeer[0] == 0 {
			x := strings.IndexRune(row, 'S')
			if x != -1 {
				reindeer = [2]int{y, x}
			}
		}
	}

	return grid, reindeer
}

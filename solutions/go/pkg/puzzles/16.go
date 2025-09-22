package puzzles

import (
	"container/heap"
	"lorech/advent-of-code-2024/pkg/grid"
	"math"
	"slices"
	"strings"
)

type tile struct {
	position  grid.Coordinates   // The location of the tile.
	direction grid.Direction     // The direction faced when traversing the tile.
	cost      int                // The cost to get to this tile; priority within the queue.
	path      []grid.Coordinates // The path leading up to the tile.
	index     int                // The index within the heap.
}

type state struct {
	position  grid.Coordinates
	direction grid.Direction
}

type priorityQueue []*tile

// Returns the length of the priority queue.
func (pq priorityQueue) Len() int { return len(pq) }

// Min-heap priority queue comparison implementation.
func (pq priorityQueue) Less(i int, j int) bool { return pq[i].cost < pq[j].cost }

// Swaps two elements around within the priority queue at the provided indices.
func (pq priorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Pushes an element into the priority queue.
func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	tile := x.(*tile)
	tile.index = n
	*pq = append(*pq, tile)
}

// Pops the final element of the priority queue.
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	tile := old[n-1]
	old[n-1] = nil  // Allow the garbage collector to reclaim the memory.
	tile.index = -1 // Prevent unintended access.
	*pq = old[:n-1]
	return tile
}

// Updates a tile within the priority queue.
func (pq *priorityQueue) update(n *tile, position grid.Coordinates, direction grid.Direction, cost int) {
	n.position = position
	n.direction = direction
	n.cost = cost
	heap.Fix(pq, n.index)
}

// Day 16: Reindeer Maze
// https://adventofcode.com/2024/day/16
func daySixteen(input string) (int, int) {
	return d16p1(input), d16p2(input)
}

// Completes the first half of the puzzle for day 16.
func d16p1(input string) int {
	maze, start, end := parseMaze(input)
	score, _ := navigate(maze, start, end)
	return score
}

// Completes the second half of the puzzle for day 16.
func d16p2(input string) int {
	maze, start, end := parseMaze(input)
	_, paths := navigate(maze, start, end)
	seats := make([]grid.Coordinates, 0)
	for _, path := range paths {
		for _, p := range path {
			if !slices.Contains(seats, p) {
				seats = append(seats, p)
			}
		}
	}
	return len(seats)
}

// Navigates the maze using Dijkstra's algorithm, returning the cheapest path
// from the start tile to the end tile, and a slice of coordinates representing
// the obtained path. Returns -1 and nil if no path is found.
func navigate(maze [][]rune, start [2]int, end [2]int) (int, [][]grid.Coordinates) {
	dirs := [4]grid.Direction{grid.Up, grid.Down, grid.Left, grid.Right}
	pq := priorityQueue{
		&tile{
			position: grid.Coordinates{
				X: start[1],
				Y: start[0],
			},
			direction: grid.Right,
			cost:      0,
			path: []grid.Coordinates{
				{
					X: start[1],
					Y: start[0],
				},
			},
		},
	}
	heap.Init(&pq)

	visited := make(map[state]int, len(maze)) // Visited tiles with their respective costs.
	var paths [][]grid.Coordinates            // Paths leading to the end of the maze.
	minCost := math.MaxInt

	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*tile)
		cost, direction, coords, path := n.cost, n.direction, n.position, n.path
		st := state{coords, direction}

		// Skip this tile if we've been here before with a lower or equal price.
		if prevCost, found := visited[st]; found && cost > prevCost {
			continue
		}
		visited[st] = cost

		// Check if we reached the end, at which point we can return!
		if coords.X == end[1] && coords.Y == end[0] {
			if cost < minCost {
				paths = [][]grid.Coordinates{path}
				minCost = cost
			} else if cost == minCost {
				paths = append(paths, path)
			}
			continue
		}

		for _, dir := range dirs {
			yd, xd := dir.Velocity()
			y, x := coords.Y+yd, coords.X+xd
			newCoords := grid.Coordinates{X: x, Y: y}

			if maze[y][x] == '#' {
				continue
			}

			c := 1 // Default value of 1 for going straight.
			if direction == dir {
				// Do nothing when going straight - the default value is already set.
			} else if (direction == grid.Up && dir == grid.Down) ||
				(direction == grid.Down && dir == grid.Up) ||
				(direction == grid.Left && dir == grid.Right) ||
				(direction == grid.Right && dir == grid.Left) {
				c += 2000 // Turn 180 degrees.
			} else {
				c += 1000 // Turn 90 degrees.
			}

			st := state{newCoords, dir}
			newCost := cost + c

			// Add to queue if not visited before or if it's cheaper this time around.
			if prevCost, found := visited[st]; !found || newCost < prevCost {
				p := make([]grid.Coordinates, len(path)+1)
				copy(p, path)
				p[len(path)] = newCoords
				heap.Push(&pq, &tile{
					position:  newCoords,
					direction: dir,
					cost:      newCost,
					path:      p,
				})
			}
		}
	}

	return minCost, paths
}

// Parses the input data into a 2D slice of runes representing the maze, and
// returns it along with the starting position of the reindeer in the maze,
// and the finish line of the maze.
func parseMaze(input string) ([][]rune, [2]int, [2]int) {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows))
	var reindeer [2]int
	var finish [2]int

	for y, row := range rows {
		grid[y] = make([]rune, len(row))
		grid[y] = []rune(row)

		if reindeer[0] == 0 {
			x := strings.IndexRune(row, 'S')
			if x != -1 {
				reindeer = [2]int{y, x}
			}
		}

		if finish[0] == 0 {
			x := strings.IndexRune(row, 'E')
			if x != -1 {
				finish = [2]int{y, x}
			}
		}
	}

	return grid, reindeer, finish
}

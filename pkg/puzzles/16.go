package puzzles

import (
	"container/heap"
	"lorech/advent-of-code-2024/pkg/grid"
	"strings"
)

type tile struct {
	position  grid.Coordinates // The location of the tile.
	direction grid.Direction   // The direction faced when traversing the tile.
	priority  int              // The priority of the tile within the queue.
	index     int              // The index within the heap.
}

type priorityQueue []*tile

// Returns the length of the priority queue.
func (pq priorityQueue) Len() int { return len(pq) }

// Compares the priority of the elements within the provided indices in the
// priority queue. As this implements a min-heap, the first element should be
// smaller than the second element.
func (pq priorityQueue) Less(i int, j int) bool {
	return pq[i].priority < pq[j].priority
}

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
func (pq *priorityQueue) update(n *tile, position grid.Coordinates, direction grid.Direction, priority int) {
	n.position = position
	n.direction = direction
	n.priority = priority
	heap.Fix(pq, n.index)
}

// Day 16: Reindeer Maze
// https://adventofcode.com/2024/day/16
func daySixteen(input string) (int, int) {
	return d16p1(input), 0
}

// Completes the first half of the puzzle for day 16.
func d16p1(input string) int {
	maze, start, end := parseMaze(input)
	visited := make([][]bool, len(maze))
	for y := range visited {
		visited[y] = make([]bool, len(maze[0]))
	}
	score := navigate(maze, start, end)
	return score
}

// Navigates the maze using Dijkstra's algorithm, returning the cheapest path
// from the start tile to the end tile. Returns -1 if no path is found.
func navigate(maze [][]rune, start [2]int, end [2]int) int {
	dirs := [4]grid.Direction{grid.Up, grid.Down, grid.Left, grid.Right}
	pq := priorityQueue{
		&tile{
			position: grid.Coordinates{
				X: start[1],
				Y: start[0],
			},
			direction: grid.Right,
			priority:  0,
		},
	}
	heap.Init(&pq)

	// Track visited tiles to prevent backtracking.
	visited := make([][]map[grid.Direction]bool, len(maze))
	for y := range visited {
		visited[y] = make([]map[grid.Direction]bool, len(maze[0]))
		for x := range visited[y] {
			visited[y][x] = make(map[grid.Direction]bool)
		}
	}

	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*tile)
		cost, direction, coords := n.priority, n.direction, n.position

		// Prevent backtracking onto this tile in the same direction.
		if visited[coords.Y][coords.X][direction] {
			continue
		}
		visited[coords.Y][coords.X][direction] = true

		// Check if we reached the end, at which point we can return the final cost!
		if coords.X == end[1] && coords.Y == end[0] {
			return cost
		}

		for _, dir := range dirs {
			yd, xd := dir.Velocity()
			y, x := coords.Y+yd, coords.X+xd

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

			nn := &tile{
				position:  grid.Coordinates{X: x, Y: y},
				direction: dir,
				priority:  cost + c,
			}

			// Add to queue if not visited.
			if !visited[y][x][dir] {
				heap.Push(&pq, nn)
			}
		}
	}

	return -1
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

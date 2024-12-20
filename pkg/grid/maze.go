package grid

// Navigates the maze using BFS, returning the finishing node, whose parent
// nodes, chained all the way to the top, will form the shortest path to reach
// the end. A second return value is provided as an indicator if the end was
// actually reached.
func NavigateMaze(maze [][]rune, start Coordinates, end Coordinates) (Tile, bool) {
	q := make([]Tile, 0)
	v := make([][]bool, len(maze))
	for i := range v {
		v[i] = make([]bool, len(maze[0]))
	}

	q = append(q, Tile{
		Position: start,
		Value:    maze[start.Y][start.X],
		Parent:   nil,
	})
	v[0][0] = true

	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		if t.Position.Y == end.Y && t.Position.X == end.X {
			return t, true
		}

		dirs := [4]Direction{Up, Down, Left, Right}
		for _, dir := range dirs {
			yd, xd := dir.Velocity()
			ny, nx := t.Position.Y+yd, t.Position.X+xd
			if ny >= 0 && ny < len(maze) && nx >= 0 && nx < len(maze[0]) && !v[ny][nx] && maze[ny][nx] != '#' {
				v[ny][nx] = true
				q = append(q, Tile{
					Position: Coordinates{X: nx, Y: ny},
					Value:    maze[ny][nx],
					Parent:   &t,
				})
			}
		}
	}

	return Tile{}, false
}

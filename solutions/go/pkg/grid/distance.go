package grid

// Calculates the Manhattan distance between a pair of coordinates.
// https://en.wikipedia.org/wiki/Taxicab_geometry
func ManhattanDistance(a Coordinate, b Coordinate) int {
	xd := a.X - b.X
	xd = max(xd, -xd)
	yd := a.Y - b.Y
	yd = max(yd, -yd)
	return xd + yd
}

// Calculates all the coordinates that are d Manhattan distance away from the
// coordinate c.
func AwayByManhattanDistance(c Coordinate, d int) []Coordinate {
	coords := make([]Coordinate, 0)
	for o := range d {
		io := d - o
		coords = append(coords, Coordinate{X: c.X + o, Y: c.Y + io})
		coords = append(coords, Coordinate{X: c.X + io, Y: c.Y - o})
		coords = append(coords, Coordinate{X: c.X - o, Y: c.Y - io})
		coords = append(coords, Coordinate{X: c.X - io, Y: c.Y + o})
	}
	return coords
}

// Calculates all the coordinates that are no more than d Manhattan distance
// away from coordinate c.
func WithinManhattanDistance(c Coordinate, d int) []Coordinate {
	coords := make([]Coordinate, 0)
	for x := -d; x <= d; x++ {
		for y := -d; y <= d; y++ {
			if max(x, -x)+max(y, -y) <= d {
				coords = append(coords, Coordinate{X: c.X + x, Y: c.Y + y})
			}
		}
	}
	return coords
}

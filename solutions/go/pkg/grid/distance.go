package grid

// Calculates the Manhattan distance between a pair of coordinates.
// https://en.wikipedia.org/wiki/Taxicab_geometry
func ManhattanDistance(a Coordinates, b Coordinates) int {
	xd := a.X - b.X
	xd = max(xd, -xd)
	yd := a.Y - b.Y
	yd = max(yd, -yd)
	return xd + yd
}

// Calculates all the coordinates that are d Manhattan distance away from the
// coordinate c.
func AwayByManhattanDistance(c Coordinates, d int) []Coordinates {
	coords := make([]Coordinates, 0)
	for o := range d {
		io := d - o
		coords = append(coords, Coordinates{X: c.X + o, Y: c.Y + io})
		coords = append(coords, Coordinates{X: c.X + io, Y: c.Y - o})
		coords = append(coords, Coordinates{X: c.X - o, Y: c.Y - io})
		coords = append(coords, Coordinates{X: c.X - io, Y: c.Y + o})
	}
	return coords
}

// Calculates all the coordinates that are no more than d Manhattan distance
// away from coordinate c.
func WithinManhattanDistance(c Coordinates, d int) []Coordinates {
	coords := make([]Coordinates, 0)
	for x := -d; x <= d; x++ {
		for y := -d; y <= d; y++ {
			if max(x, -x)+max(y, -y) <= d {
				coords = append(coords, Coordinates{X: c.X + x, Y: c.Y + y})
			}
		}
	}
	return coords
}

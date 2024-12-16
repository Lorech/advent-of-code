package grid

type Direction int

// An enum, containing pointing directions in a 2-axis grid.
const (
	Up Direction = iota
	Down
	Left
	Right
)

// Rotates the provided direction clockwise by one step.
func (p *Direction) Clockwise() {
	switch *p {
	case Up:
		*p = Right
	case Right:
		*p = Down
	case Down:
		*p = Left
	case Left:
		*p = Up
	}
}

// Rotates the provided direction counter-clockwise by one step.
func (p *Direction) CounterClockwise() {
	switch *p {
	case Up:
		*p = Left
	case Right:
		*p = Up
	case Down:
		*p = Right
	case Left:
		*p = Down
	}
}

// Determines the amount of y and x tiles to move based on the current direction.
func (p Direction) Velocity() (int, int) {
	var (
		yd, xd int
	)

	switch p {
	case Up:
		yd = -1
		xd = 0
	case Down:
		yd = 1
		xd = 0
	case Left:
		yd = 0
		xd = -1
	case Right:
		yd = 0
		xd = 1
	}

	return yd, xd
}

package grid

import (
	"testing"
)

// Tests Manhattan distance calculation.
func TestManhattanDistance(t *testing.T) {
	a, b := Coordinates{X: 0, Y: 0}, Coordinates{X: 6, Y: 6}
	e := 12
	if r := ManhattanDistance(a, b); r != e {
		t.Errorf("ManhattanDistance = %v, expected %v", r, e)
	}
}

// Tests Manhattan distance calculation where A is larger than B.
func TestInverseManhattanDistance(t *testing.T) {
	a, b := Coordinates{X: 6, Y: 6}, Coordinates{X: 0, Y: 0}
	e := 12
	if r := ManhattanDistance(a, b); r != e {
		t.Errorf("ManhattanDistance = %v, expected %v", r, e)
	}
}

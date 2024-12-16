package grid

import "testing"

// Tests clockwise rotation of a direction.
func TestClockwise(t *testing.T) {
	r := Up
	r.Clockwise()
	e := Right
	if r != e {
		t.Errorf("Clockwise() = %v, expected %v", r, e)
	}
}

// Tests counter-clockwise rotation of a direction.
func TestCounterClockwise(t *testing.T) {
	r := Up
	r.CounterClockwise()
	e := Left
	if r != e {
		t.Errorf("CounterClockwise() = %v, expected %v", r, e)
	}
}

// Tests obtaining the velocity of a direction.
func TestVelocity(t *testing.T) {
	p := Right
	r1, r2 := p.Velocity()
	e1, e2 := 0, 1
	if r1 != e1 || r2 != e2 {
		t.Errorf("Velocity = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

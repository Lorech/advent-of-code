package grid

import "testing"

// Tests clockwise rotation of a direction.
func TestRotate(t *testing.T) {
	r := Up
	r.Rotate()
	e := Right
	if r != e {
		t.Errorf("Rotate() = %v, expected %v", r, e)
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

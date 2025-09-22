package convert

import "testing"

// Tests slice conversion to integer with default configuration.
func TestSliceToInt(t *testing.T) {
	s := []int{1, 2, 3}
	e := 123
	if r, err := Stoi(s); err != nil || r != e {
		t.Errorf("Stoi() = %v, expected %v", r, e)
	}
}

// Tests slice conversion to integer with a different base.
func TestSliceToIntDifferentBase(t *testing.T) {
	s := []int{1, 2, 3}
	e := 83
	if r, err := Stoi(s, 8); err != nil || r != e {
		t.Errorf("Stoi() = %v, expected %v", r, e)
	}
}

// Tests slice conversion to integer with leading zeroes.
func TestSliceToIntWithZeroes(t *testing.T) {
	s := []int{0, 0, 0, 1, 2, 3}
	e := 123
	if r, err := Stoi(s); err != nil || r != e {
		t.Errorf("Stoi() = %v, expected %v", r, e)
	}
}

// Tests slice conversion to integer with an invalid slice.
func TestInvalidSliceToInt(t *testing.T) {
	s := []int{1, 2, 0}
	e := -1
	if r, err := Stoi(s, 1); err == nil || r != e {
		t.Errorf("Stoi() = %v, expected %v", r, e)
	}
}

package convert

import "testing"

// Tests converting a true boolean to an integer.
func TestBtoiTrue(t *testing.T) {
	if r := Btoi(true); r != 1 {
		t.Errorf("Btoi() = %v, expected %v", r, 1)
	}
}

// Tests converting a false boolean to an integer.
func TestBtoiFalse(t *testing.T) {
	if r := Btoi(false); r != 0 {
		t.Errorf("Btoi() = %v, expected %v", r, 0)
	}
}

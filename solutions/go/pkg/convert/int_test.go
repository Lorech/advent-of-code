package convert

import "testing"

// Tests converting an int 0 to a false boolean.
func TestItobFalse(t *testing.T) {
	if r := Itob(0); r != false {
		t.Errorf("Itob() = %v, expected %v", r, false)
	}
}

// Tests converting a positive int to a true boolean.
func TestItobPositive(t *testing.T) {
	if r := Itob(1); r != true {
		t.Errorf("Itob() = %v, expected %v", r, true)
	}
}

// Tests converting a negative int to a true boolean.
func TestItobNegative(t *testing.T) {
	if r := Itob(-1); r != true {
		t.Errorf("Itob() = %v, expected %v", r, true)
	}
}

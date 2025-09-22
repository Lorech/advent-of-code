package cslices

import "testing"

// Tests the removal of the first element from an integer slice by index.
func TestRemoveFirstIntByIndex(t *testing.T) {
	s := []int{1, 2, 3}
	e := []int{2, 3}
	r := Remove(s, 0)

	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("Remove() = %v, expected %v", r, e)
	}
}

// Tests the removal of a middle element from an integer slice by index.
func TestRemoveMiddleIntByIndex(t *testing.T) {
	s := []int{1, 2, 3}
	e := []int{1, 3}
	r := Remove(s, 1)

	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("Remove() = %v, expected %v", r, e)
	}
}

// Tests the removal of the last element from an integer slice by index.
func TestRemoveLastIntByIndex(t *testing.T) {
	s := []int{1, 2, 3}
	e := []int{1, 2}
	r := Remove(s, 2)

	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("Remove() = %v, expected %v", r, e)
	}
}

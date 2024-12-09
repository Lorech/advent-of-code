package utils

import "testing"

// Tests the removal of the first element from an integer slice by index.
func TestRemoveFirstIntByIndex(t *testing.T) {
	s := []int{1, 2, 3}
	e := []int{2, 3}
	r := RemoveInt(s, 0)

	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("RemoveInt() = %v, expected %v", r, e)
	}
}

// Tests the removal of a middle element from an integer slice by index.
func TestRemoveMiddleIntByIndex(t *testing.T) {
	s := []int{1, 2, 3}
	e := []int{1, 3}
	r := RemoveInt(s, 1)

	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("RemoveInt() = %v, expected %v", r, e)
	}
}

// Tests the removal of the last element from an integer slice by index.
func TestRemoveLastIntByIndex(t *testing.T) {
	s := []int{1, 2, 3}
	e := []int{1, 2}
	r := RemoveInt(s, 2)

	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("RemoveInt() = %v, expected %v", r, e)
	}
}

// Tests subslice index retrieval with both values in the slice.
func TestValidSubsliceIndex(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	e1, e2 := 1, 3
	r1, r2 := SubsliceIndex(s, 2, 3)
	if r1 != e1 || r2 != e2 {
		t.Errorf("SubsliceIndex() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

// Tests subslice retrieval with both values in the slice.
func TestValidSubslice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	e := []int{2, 3}
	r := Subslice(s, 2, 3)
	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("Subslice() = %v, expected %v", r, e)
	}
}

// Tests subslice index retrieval with only the second value in the slice.
func TestEndingSubsliceIndex(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	e1, e2 := 0, 3
	r1, r2 := SubsliceIndex(s, 6, 3)
	if r1 != e1 || r2 != e2 {
		t.Errorf("SubsliceIndex() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

// Tests subslice retrieval with only the second value in the slice.
func TestEndingSubslice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	e := []int{1, 2, 3}
	r := Subslice(s, 6, 3)
	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] || r[2] != e[2] {
		t.Errorf("Subslice() = %v, expected %v", r, e)
	}
}

// Tests subslice index retrieval with only the first value in the slice.
func TestStartingSubsliceIndex(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	e1, e2 := 1, 5
	r1, r2 := SubsliceIndex(s, 2, 6)
	if r1 != e1 || r2 != e2 {
		t.Errorf("SubsliceIndex() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

// Tests subslice retrieval with only the first value in the slice.
func TestStartingSubslice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	e := []int{2, 3, 4, 5}
	r := Subslice(s, 2, 6)
	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] || r[2] != e[2] || r[3] != e[3] {
		t.Errorf("Subslice() = %v, expected %v", r, e)
	}
}

// Tests subslice index retrieval with neither value in the slice.
func TestInvalidSubsliceIndex(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	e1, e2 := 0, 5
	r1, r2 := SubsliceIndex(s, 6, 7)
	if r1 != e1 || r2 != e2 {
		t.Errorf("SubsliceIndex() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

// Tests subslice retrieval with neither value in the slice.
func TestInvalidSubslice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	r := Subslice(s, 6, 7)
	if len(r) != len(s) || r[0] != s[0] || r[1] != s[1] || r[2] != s[2] || r[3] != s[3] || r[4] != s[4] {
		t.Errorf("Subslice() = %v, expected %v", r, s)
	}
}

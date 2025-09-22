package cslices

import "testing"

// Tests the appearance of a single sequence within a slice.
func TestAppearsOneTime(t *testing.T) {
	s := []int{1, 2, 2, 2, 3}
	e1 := 1
	e2 := 4
	r1, r2 := Appears(s, 2)

	if r1 != e1 || r2 != e2 {
		t.Errorf("Appears() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

// Tests the appearance of a multiple sequences within a slice.
func TestAppearsMultipleTimes(t *testing.T) {
	s := []int{1, 2, 2, 2, 3, 2, 2, 2}
	e1 := 1
	e2 := 4
	r1, r2 := Appears(s, 2)

	if r1 != e1 || r2 != e2 {
		t.Errorf("Appears() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

// Tests the absence of a sequence within a slice.
func TestAppearsNoTimes(t *testing.T) {
	s := []int{1, 3, 4, 5, 6}
	e1 := -1
	e2 := -1
	r1, r2 := Appears(s, 2)

	if r1 != e1 || r2 != e2 {
		t.Errorf("Appears() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

// Tests the overall appearances of a single sequence within a slice.
func TestAppearsAllOneTime(t *testing.T) {
	s := []int{1, 2, 2, 2, 3}
	e := [][2]int{{1, 4}}
	r := AppearsAll(s, 2)

	if len(r) != len(e) || r[0] != e[0] {
		t.Errorf("AppearsAll() = %v, expected %v", r, e)
	}
}

// Tests the overall appearance of multiple sequences within a slice.
func TestAppearsAllMultipleTimes(t *testing.T) {
	s := []int{1, 2, 2, 2, 3, 2, 2, 2}
	e := [][2]int{{1, 4}, {5, 8}}
	r := AppearsAll(s, 2)

	if len(r) != len(e) || r[0] != e[0] || r[1] != e[1] {
		t.Errorf("AppearsAll() = %v, expected %v", r, e)
	}
}

// Tests the overall absence of a sequence within a slice.
func TestAppearsAllNoTimes(t *testing.T) {
	s := []int{1, 3, 4, 5, 6}
	e := [][2]int{}
	r := AppearsAll(s, 2)

	if len(r) != len(e) {
		t.Errorf("AppearsAll() = %v, expected %v", r, e)
	}
}

package cslices

import (
	"slices"
)

// Finds the first appearence chain of a value within a slice, returning the
// start and end positions of this value, or -1 if the value does not appear.
//
// The end position is non-inclusive to match slicing behavior when creating
// subslices.
func Appears[S ~[]V, V comparable](s S, v V) (int, int) {
	start := slices.Index(s, v)

	if start == -1 {
		return -1, -1
	}

	end := len(s)
	for i := start + 1; i < len(s); i++ {
		if s[i] != v {
			end = i
			break
		}
	}

	return start, end
}

// Find all appearence chains of a value within a slice, returning a new slice
// containing arrays of all start and end positions of this value.
func AppearsAll[S ~[]V, V comparable](s S, v V) [][2]int {
	r := make([][2]int, 0)
	start := 0

	for {
		a, b := Appears(s[start:], v)

		if a == -1 || b == -1 {
			break
		}

		r = append(r, [2]int{a + start, b + start})
		start += b
	}

	return r
}

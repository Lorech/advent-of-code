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

	end := start + 1
	for i := end; i < len(s); i++ {
		if s[i] != v {
			end = i
			break
		}
	}

	return start, end
}

package cslices

import "slices"

// Get the indexes of the first of a and its' closest occurance of b within
// slice s.
//
// If either value is not found in the slice, the respective boundrary of the
// value will match the respective end of the original slice.
func SubsliceIndex[S ~[]E, E comparable](s S, a E, b E) (int, int) {
	// The start is the position of a or the start of the original slice.
	start := slices.Index(s, a)
	if start == -1 {
		start = 0
	}

	// The end is the end of the original slice or the closest position of b to a.
	end := len(s)
	for i := start + 1; i != len(s); i++ {
		if s[i] == b {
			// i is off-by-one because slicing uses the second value non-inclusively.
			end = i + 1
			break
		}
	}

	return start, end
}

// Get a subslice from the first occurance of a to its' closest occurance of b
// within slice s.
//
// If either value is not found in the slice, the respective boundrary of the
// value will match the respective end of the original slice.
func Subslice[S ~[]E, E comparable](s S, a E, b E) []E {
	start, end := SubsliceIndex(s, a, b)
	return s[start:end]
}

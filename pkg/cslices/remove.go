package cslices

// Remove an element from a slice by index.
//
// If the index is out of bounds, the original slice is returned.
func Remove[S ~[]I, I any](s S, i int) S {
	if i < 0 || i >= len(s) {
		return s
	}
	result := append([]I{}, s...) // Copy the slice to avoid modifying the original.
	return append(result[:i], result[i+1:]...)
}

package utils

// Remove an element from a slice by index.
//
// If the index is out of bounds, the original slice is returned.
func Remove(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result := append([]int{}, slice...) // Copy the slice to avoid modifying the original.
	return append(result[:index], result[index+1:]...)
}

package convert

// Converts a boolean to an integer.
// true == 1, false == 0
func Btoi(b bool) int {
	if b {
		return 1
	}

	return 0
}

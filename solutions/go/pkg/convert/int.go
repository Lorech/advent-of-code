package convert

// Converts an integer to a boolean representation of it.
// 0 == false, any other value == true.
func Itob(i int) bool {
	if i == 0 {
		return false
	}

	return true
}

package convert

import "slices"

// Converts an integer to a boolean representation of it.
// 0 == false, any other value == true.
func Itob(i int) bool {
	if i == 0 {
		return false
	}

	return true
}

// Converts a binary number into a slice of integers with each number
// within the number representing an index from MSB to LSB that contains a 1.
func BinToIntIndex(num, digits int) []int {
	b, i := make([]int, digits), digits-1
	for num != 0 {
		b[i] = num & 0x1
		num >>= 1
		i--
	}

	r, p := make([]int, 0), 0
	for true {
		i := slices.Index(b, 1)
		if i == -1 {
			break
		}
		r = append(r, i+p)
		b = b[i+1:]
		p += i + 1
	}
	return r
}

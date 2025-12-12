package convert

import (
	"fmt"
	"math"
)

// Converts a slice of integers to a base 10 integer. The slice is expected to
// be ordered from most to least significant digit. The optional second param
// allows specifying the base to convert from, which defaults to 10. Returns an
// error if the input is invalid due to providing larger input values than base.
func Stoi(s []int, b ...int) (int, error) {
	base := 10
	if len(b) > 0 {
		base = b[0]
	}

	// Strip leading zeroes.
	for i, d := range s {
		if d != 0 {
			s = s[i:]
			break
		}
	}

	num := 0
	for i, d := range s {
		if d >= base {
			return -1, fmt.Errorf("%d does not fit in base %d\n", d, base)
		}

		num += int(math.Pow(float64(base), float64(len(s)-1-i))) * d
	}
	return num, nil
}

// Converts an integer slice into a single binary number with each number
// within the slice representing an index from MSB to LSB that contains a 1.
func IntIndexToBinary(num []int) int {
	r, p := 0, 0
	for _, n := range num {
		r <<= n - p
		r |= 0x1
		p = n
	}
	return r
}

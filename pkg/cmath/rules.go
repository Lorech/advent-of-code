package cmath

// Finds the greatest common divisor for two integers.
//
// Based on the Euclidian algorithm. See https://en.wikipedia.org/wiki/Euclidean_algorithm
func Gcd(a int, b int) int {
	if b == 0 {
		return a
	}

	return Gcd(b, a%b)
}

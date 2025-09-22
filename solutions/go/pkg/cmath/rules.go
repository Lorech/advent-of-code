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

// Finds the median in a slice of number values.
//
// The function assumes that the provided slice is already sorted.
func Median[V int | int32 | int64 | float32 | float64](s []V) float64 {
	if len(s) == 0 {
		return float64(0)
	}

	if len(s) == 1 {
		return float64(s[0])
	}

	if len(s)%2 == 0 {
		b := len(s) / 2
		a := b - 1
		return (float64(s[a]) + float64(s[b])) / 2
	} else {
		return float64(s[(len(s)-1)/2])
	}
}

package cmath

// Solves a system of two linear equations given in the format:
//   - a1 * x + b1 * y = r1
//   - a2 * x + b2 * y = r2
//
// The return value is a tuple containing the values of x and y. If both values
// equal zero, there was no integer solution possible to solve the system.
func CramersRule(a1 int, b1 int, r1 int, a2 int, b2 int, r2 int) (int, int) {
	divisor := a1*b2 - a2*b1
	aDividend := r1*b2 - r2*b1
	bDividend := a1*r2 - a2*r1

	// There is no integer solution for this system.
	if aDividend%divisor != 0 || bDividend%divisor != 0 {
		return 0, 0
	}

	return aDividend / divisor, bDividend / divisor
}

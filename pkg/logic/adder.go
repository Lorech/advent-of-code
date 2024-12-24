package logic

import (
	"lorech/advent-of-code-2024/pkg/convert"
)

// Implements a half adder, which sums two bits of data, returning the sum and
// the carry bit of the operation.
func HalfAdder(a bool, b bool) (bool, bool) {
	sum := a != b
	carry := a && b
	return sum, carry
}

// Implements a full adder, which sums two bits of data with respect to a carry
// bit, returning the sum and the new carry bit of the operation.
func FullAdder(a bool, b bool, c bool) (bool, bool) {
	sum, carry := HalfAdder(a, b)
	c1 := c && sum
	sum = sum != c
	carry = convert.Itob(convert.Btoi(carry) + convert.Btoi(c1))
	return sum, carry
}

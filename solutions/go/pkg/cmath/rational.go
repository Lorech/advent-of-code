package cmath

// A rational number in canonical form to allow integer-only calculation.
type Rational struct {
	Numerator   int
	Denominator int
}

// Creates a new rational number with numerator `n` and denominator `d`.
// Automatically reduces n and d to their minimum values while retaining the fraction.
func NewRational(n, d int) Rational {
	r := Rational{Numerator: n, Denominator: d}
	r.reduce()
	return r
}

// Convenience method for converting an integer to canonical rational number form.
func NewRationalInteger(n int) Rational {
	// Intentionally avoid reusing `NewRational` to avoid unnecessary reduction
	return Rational{Numerator: n, Denominator: 1}
}

// Check if two rational numbers are equal.
// This is a strict equality check of the number's parts, meaning that an equal fraction
// will not necessarily indicate an equal rational number.
func (a Rational) Equal(b Rational) bool {
	return a.Numerator == b.Numerator && a.Denominator == b.Denominator
}

// Check if a rational number is an integer.
func (a Rational) Integer() bool {
	return a.Denominator == 1
}

// Adds two rational numbers, returning the resulting rational number.
func (a Rational) Add(b Rational) Rational {
	return NewRational(
		a.Numerator*b.Denominator+a.Denominator*b.Numerator,
		a.Denominator*b.Denominator,
	)
}

// Subtracts two rational numbers, returning the resulting rational number.
func (a Rational) Sub(b Rational) Rational {
	return NewRational(
		a.Numerator*b.Denominator-a.Denominator*b.Numerator,
		a.Denominator*b.Denominator,
	)
}

// Multiplies two rational numbers, returning the resulting rational number.
func (a Rational) Mul(b Rational) Rational {
	return NewRational(
		a.Numerator*b.Numerator,
		a.Denominator*b.Denominator,
	)
}

// Divides two rational numbers, returning the resulting rational number.
func (a Rational) Div(b Rational) Rational {
	r := Rational{Numerator: b.Denominator, Denominator: b.Numerator}
	return a.Mul(r)
}

// Reduces the fraction of the given rational number using its GCD.
func (r *Rational) reduce() {
	d := Gcd(r.Numerator, r.Denominator)
	r.Numerator /= d
	r.Denominator /= d
}

package cmath

import "testing"

// Tests obtaining the greatest common divisor for numbers that can be
// successfully divided by a value larger than 1.
func TestGcd(t *testing.T) {
	e := 4
	if r := Gcd(36, 28); e != r {
		t.Errorf("Gcd() = %v, expected %v", r, e)
	}
}

// Tests obtaining the greatest common divisor for numbers that can be
// successfully divided by a value larger than 1, where b>a.
func TestInvertedGcd(t *testing.T) {
	e := 4
	if r := Gcd(28, 36); e != r {
		t.Errorf("Gcd() = %v, expected %v", r, e)
	}
}

// Tests obtaining the greatest common divisor for numbers that are coprime.
func TestCoprimeGcd(t *testing.T) {
	e := 1
	if r := Gcd(28, 9); e != r {
		t.Errorf("Gcd() = %v, expected %v", r, e)
	}
}

// Tests obtaining the greatest common divisor for numbers that are coprime,
// where b>a.
func TestInvertedCoprimeGcd(t *testing.T) {
	e := 1
	if r := Gcd(9, 28); e != r {
		t.Errorf("Gcd() = %v, expected %v", r, e)
	}
}

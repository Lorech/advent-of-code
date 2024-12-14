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

// Tests obtaining the median of an even length slice of integers.
func TestEvenIntMedian(t *testing.T) {
	e := 4.5
	if r := Median([]int{1, 2, 3, 4, 5, 6, 8, 9}); e != r {
		t.Errorf("Median() = %v, expected %v", r, e)
	}
}

// Tests obtaining the median of an odd length slice of integers.
func TestOddIntMedian(t *testing.T) {
	e := float64(5)
	if r := Median([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}); e != r {
		t.Errorf("Median() = %v, expected %v", r, e)
	}
}

// Tests obtaining the median of an even length slice of floats.
func TestEvenFloatMedian(t *testing.T) {
	e := 1.5
	if r := Median([]float64{1, 1.2, 1.4, 1.6, 1.8, 2}); e != r {
		t.Errorf("Median() = %v, expected %v", r, e)
	}
}

// Tests obtaining the median of an odd length slice of floats.
func TestOddFloatMedian(t *testing.T) {
	e := 1.4
	if r := Median([]float64{1, 1.2, 1.4, 1.8, 2}); e != r {
		t.Errorf("Median() = %v, expected %v", r, e)
	}
}

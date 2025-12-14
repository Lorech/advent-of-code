package cmath_test

import (
	"lorech/advent-of-code/pkg/cmath"
	"testing"
)

func TestNewRational(t *testing.T) {
	tests := []struct {
		// Description of test case
		name string
		// Input parameters for target function
		n    int
		d    int
		want cmath.Rational
	}{
		{
			name: "reduced fraction",
			n:    1,
			d:    2,
			want: cmath.Rational{Numerator: 1, Denominator: 2},
		},
		{
			name: "reducable fraction",
			n:    8,
			d:    16,
			want: cmath.Rational{Numerator: 1, Denominator: 2},
		},
		{
			name: "irreducable fraction",
			n:    5,
			d:    12,
			want: cmath.Rational{Numerator: 5, Denominator: 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmath.NewRational(tt.n, tt.d)
			if tt.want.Numerator != got.Numerator || tt.want.Denominator != got.Denominator {
				t.Errorf("NewRational() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRationalInteger(t *testing.T) {
	tests := []struct {
		// Description of test case
		name string
		// Input parameters for target function
		n    int
		want cmath.Rational
	}{
		{
			name: "rational integer",
			n:    5,
			want: cmath.Rational{Numerator: 5, Denominator: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmath.NewRationalInteger(tt.n)
			if tt.want.Numerator != got.Numerator || tt.want.Denominator != got.Denominator {
				t.Errorf("NewRationalInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRational_Equal(t *testing.T) {
	tests := []struct {
		// Description of test case
		name string
		// Input parameters for receiver constructor
		n int
		d int
		// Input parameters for target function
		b    cmath.Rational
		want bool
	}{
		{
			name: "equal rational",
			n:    1,
			d:    2,
			b:    cmath.Rational{Numerator: 1, Denominator: 2},
			want: true,
		},
		{
			name: "inequal rational",
			n:    1,
			d:    2,
			b:    cmath.Rational{Numerator: 1, Denominator: 3},
			want: false,
		},
		{
			name: "equal fraction",
			n:    1,
			d:    2,
			b:    cmath.Rational{Numerator: 2, Denominator: 4},
			want: false,
		},
		{
			name: "inequal fraction",
			n:    1,
			d:    2,
			b:    cmath.Rational{Numerator: 2, Denominator: 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := cmath.NewRational(tt.n, tt.d)
			got := a.Equal(tt.b)
			if got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRational_Integer(t *testing.T) {
	tests := []struct {
		// Description of test case
		name string
		// Input parameters for receiver constructor
		n    int
		d    int
		want bool
	}{
		{
			name: "integer",
			n:    5,
			d:    1,
			want: true,
		},
		{
			name: "fraction",
			n:    5,
			d:    12,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := cmath.NewRational(tt.n, tt.d)
			got := a.Integer()
			if got != tt.want {
				t.Errorf("Integer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRational_Add(t *testing.T) {
	tests := []struct {
		// Description of test case
		name string
		// Input parameters for receiver constructor
		n int
		d int
		// Input parameters for target function
		b    cmath.Rational
		want cmath.Rational
	}{
		{
			name: "common denominator",
			n:    1,
			d:    3,
			b:    cmath.Rational{Numerator: 1, Denominator: 3},
			want: cmath.Rational{Numerator: 2, Denominator: 3},
		},
		{
			name: "uncommon denominator",
			n:    1,
			d:    3,
			b:    cmath.Rational{Numerator: 2, Denominator: 5},
			want: cmath.Rational{Numerator: 11, Denominator: 15},
		},
		{
			name: "reduces to integer",
			n:    1,
			d:    3,
			b:    cmath.Rational{Numerator: 2, Denominator: 3},
			want: cmath.Rational{Numerator: 1, Denominator: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := cmath.NewRational(tt.n, tt.d)
			got := a.Add(tt.b)
			if tt.want.Numerator != got.Numerator || tt.want.Denominator != got.Denominator {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRational_Sub(t *testing.T) {
	tests := []struct {
		// Description of test case
		name string
		// Input parameters for receiver constructor
		n int
		d int
		// Input parameters for target function
		b    cmath.Rational
		want cmath.Rational
	}{
		{
			name: "common denominator",
			n:    2,
			d:    3,
			b:    cmath.Rational{Numerator: 1, Denominator: 3},
			want: cmath.Rational{Numerator: 1, Denominator: 3},
		},
		{
			name: "uncommon denominator",
			n:    2,
			d:    3,
			b:    cmath.Rational{Numerator: 2, Denominator: 5},
			want: cmath.Rational{Numerator: 4, Denominator: 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := cmath.NewRational(tt.n, tt.d)
			got := a.Sub(tt.b)
			if tt.want.Numerator != got.Numerator || tt.want.Denominator != got.Denominator {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRational_Mul(t *testing.T) {
	tests := []struct {
		// Description of test case
		name string
		// Input parameters for receiver constructor
		n int
		d int
		// Input parameters for target function
		b    cmath.Rational
		want cmath.Rational
	}{
		{
			name: "common denominator",
			n:    2,
			d:    3,
			b:    cmath.Rational{Numerator: 1, Denominator: 3},
			want: cmath.Rational{Numerator: 2, Denominator: 9},
		},
		{
			name: "uncommon denominator",
			n:    1,
			d:    3,
			b:    cmath.Rational{Numerator: 2, Denominator: 5},
			want: cmath.Rational{Numerator: 2, Denominator: 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := cmath.NewRational(tt.n, tt.d)
			got := a.Mul(tt.b)
			if tt.want.Numerator != got.Numerator || tt.want.Denominator != got.Denominator {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRational_Div(t *testing.T) {
	tests := []struct {
		// Description of test case
		name string
		// Input parameters for receiver constructor
		n int
		d int
		// Input parameters for target function
		b    cmath.Rational
		want cmath.Rational
	}{
		{
			name: "common denominator",
			n:    2,
			d:    3,
			b:    cmath.Rational{Numerator: 1, Denominator: 3},
			want: cmath.Rational{Numerator: 2, Denominator: 1},
		},
		{
			name: "uncommon denominator",
			n:    1,
			d:    3,
			b:    cmath.Rational{Numerator: 2, Denominator: 5},
			want: cmath.Rational{Numerator: 5, Denominator: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := cmath.NewRational(tt.n, tt.d)
			got := a.Div(tt.b)
			if tt.want.Numerator != got.Numerator || tt.want.Denominator != got.Denominator {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

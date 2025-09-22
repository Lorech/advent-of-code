package cmath

import "testing"

// Tests a system of linear equations with valid integer answers.
func TestValidCramersRule(t *testing.T) {
	e1, e2 := 3, -7
	if r1, r2 := CramersRule(4, 1, 5, 1, -1, 10); r1 != e1 || r2 != e2 {
		t.Errorf("CramersRule() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

// Tests a system of linear equations without valid integer answers.
func TestInvalidCramersRule(t *testing.T) {
	e1, e2 := 0, 0
	if r1, r2 := CramersRule(3, 2, -4, -1, 2, -2); r1 != e1 || r2 != e2 {
		t.Errorf("CramersRule() = %v, %v, expected %v, %v", r1, r2, e1, e2)
	}
}

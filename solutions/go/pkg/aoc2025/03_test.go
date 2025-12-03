package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 3.
func TestDayThreePartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 3)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 357
	if r := d3p1(input); e != r {
		t.Errorf("d3p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 3.
func BenchmarkDayThreePartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 3)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d3p1(input)
	}
}

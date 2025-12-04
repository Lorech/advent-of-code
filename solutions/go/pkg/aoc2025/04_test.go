package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the second part of the puzzle for day 4.
func TestDayFourPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 4)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 13
	if r := d4p1(input); e != r {
		t.Errorf("d4p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 4.
func BenchmarkDayFourPartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 4)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d4p1(input)
	}
}

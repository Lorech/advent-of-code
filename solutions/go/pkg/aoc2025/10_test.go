package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 10.
func TestDayTenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 10)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 7
	if r := d10p1(input); e != r {
		t.Errorf("d10p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 10.
func BenchmarkDayTenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 10)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d10p1(input)
	}
}

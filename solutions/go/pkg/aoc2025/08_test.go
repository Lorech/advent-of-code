package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 8
func TestDayEightPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 8)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 40
	if r := d8p1(input, 10); e != r {
		t.Errorf("d8p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 8.
func BenchmarkDayEightPartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 8)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d8p1(input)
	}
}

package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 7
func TestDaySevenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 7)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 21
	if r := d7p1(input); e != r {
		t.Errorf("d7p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 7.
func BenchmarkDaySevenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 7)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d7p1(input)
	}
}

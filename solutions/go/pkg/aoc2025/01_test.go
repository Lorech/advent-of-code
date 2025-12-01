package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 1.
func TestDayOnePartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 1)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 3
	if r := d1p1(input); e != r {
		t.Errorf("d1p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 1.
func BenchmarkDayOnePartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 1)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d1p1(input)
	}
}

// Tests the second part of the puzzle for day 1.
func TestDayOnePartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2025, 1)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 6
	if r := d1p2(input); e != r {
		t.Errorf("d1p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 1.
func BenchmarkDayOnePartTwo(b *testing.B) {
	input, err := file.ReadInfile(2025, 1)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d1p2(input)
	}
}

package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 6.
func TestDaySixPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 6)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 4277556
	if r := d6p1(input); e != r {
		t.Errorf("d6p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 6.
func BenchmarkDaySixPartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 6)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d6p1(input)
	}
}

// Tests the second part of the puzzle for day 6.
func TestDaySixPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2025, 6)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 3263827
	if r := d6p2(input); e != r {
		t.Errorf("d6p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 6.
func BenchmarkDaySixPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2025, 6)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d6p2(input)
	}
}

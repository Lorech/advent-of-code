package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 9
func TestDayNinePartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 9)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 50
	if r := d9p1(input); e != r {
		t.Errorf("d9p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 9.
func BenchmarkDayNinePartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 9)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d9p1(input)
	}
}

// Tests the second part of the puzzle for day 9
func TestDayNinePartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2025, 9)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 24
	if r := d9p2(input); e != r {
		t.Errorf("d9p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 9.
func BenchmarkDayNinePartTwo(b *testing.B) {
	input, err := file.ReadInfile(2025, 9)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d9p2(input)
	}
}

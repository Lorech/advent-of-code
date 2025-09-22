package puzzles

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 16.
func TestDaySixteenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 16)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 11048
	if r := d16p1(input); e != r {
		t.Errorf("d16p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 16.
func BenchmarkDaySixteenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 16)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d16p1(input)
	}
}

// Tests the second part of the puzzle for day 16.
func TestDaySixteenPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2024, 16)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 64
	if r := d16p2(input); e != r {
		t.Errorf("d16p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 16.
func BenchmarkDaySixteenPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 16)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d16p2(input)
	}
}

package puzzles

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 8.
func TestDayEightPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 8)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 14
	if r := d8p1(input); e != r {
		t.Errorf("d8p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 8.
func TestDayEightPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2024, 8)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 34
	if r := d8p2(input); e != r {
		t.Errorf("d8p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 8.
func BenchmarkDayEightPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 8)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d8p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 8.
func BenchmarkDayEightPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 8)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d8p2(input)
	}
}

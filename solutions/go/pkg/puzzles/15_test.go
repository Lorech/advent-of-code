package puzzles

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 15.
func TestDayFifteenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 15)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 10092
	if r := d15p1(input); e != r {
		t.Errorf("d15p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 15.
func BenchmarkDayFifteenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 15)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d15p1(input)
	}
}

// Tests the second part of the puzzle for day 15.
func TestDayFifteenPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2024, 15)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 9021
	if r := d15p2(input); e != r {
		t.Errorf("d15p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 15.
func BenchmarkDayFifteenPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 15)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d15p2(input)
	}
}

package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 4.
func TestDayFourPartOne(t *testing.T) {
	input, err := file.ReadTestFile(4)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 18
	if r := d4p1(input); e != r {
		t.Errorf("d4p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 4.
func TestDayFourPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(4)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 9
	if r := d4p2(input); e != r {
		t.Errorf("d4p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 4.
func BenchmarkDayFourPartOne(b *testing.B) {
	input, err := file.ReadInfile(4)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d4p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 4.
func BenchmarkDayFourPartTwo(b *testing.B) {
	input, err := file.ReadInfile(4)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d4p2(input)
	}
}

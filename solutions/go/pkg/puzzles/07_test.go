package puzzles

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 7.
func TestDaySevenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 7)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 3749
	if r := d7p1(input); e != r {
		t.Errorf("d7p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 7.
func TestDaySevenPartSecond(t *testing.T) {
	input, err := file.ReadTestFile(2024, 7)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 11387
	if r := d7p2(input); e != r {
		t.Errorf("d7p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 7.
func BenchmarkDaySevenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 7)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d7p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 7.
func BenchmarkDaySevenPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 7)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d7p2(input)
	}
}

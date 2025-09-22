package puzzles

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 14.
func TestDayFourteenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 14)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 12
	if r := d14p1(input, 11, 7); e != r {
		t.Errorf("d14p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 14.
func BenchmarkDayFourteenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 14)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d14p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 14.
func BenchmarkDayFourteenPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 14)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d14p2(input)
	}
}

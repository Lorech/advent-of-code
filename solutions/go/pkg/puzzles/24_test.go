package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 24.
func TestDayTwentyFourPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 24, "part1")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 2024
	if r := d24p1(input); e != r {
		t.Errorf("d24p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 24.
func BenchmarkDayTwentyFourPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 24)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d24p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 24.
func BenchmarkDayTwentyFourPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 24)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d24p2(input)
	}
}

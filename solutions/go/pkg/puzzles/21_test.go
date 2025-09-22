package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 21.
func TestDayTwentyOnePartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 21)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 126384
	if r := d21p1(input); e != r {
		t.Errorf("d21p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 21.
func BenchmarkDayTwentyOnePartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 21)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d21p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 21.
func BenchmarkDayTwentyOnePartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 21)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d21p2(input)
	}
}

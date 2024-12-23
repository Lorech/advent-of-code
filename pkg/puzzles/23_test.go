package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 23.
func TestDayTwentyThreePartOne(t *testing.T) {
	input, err := file.ReadTestFile(23)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 7
	if r := d23p1(input); e != r {
		t.Errorf("d23p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 23.
func BenchmarkDayTwentyThreePartOne(b *testing.B) {
	input, err := file.ReadInfile(23)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d23p1(input)
	}
}

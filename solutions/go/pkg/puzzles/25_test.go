package puzzles

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 25.
func TestDayTwentyFivePartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 25)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 3
	if r := d25p1(input); e != r {
		t.Errorf("d25p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 25.
func BenchmarkDayTwentyFivePartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 25)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d25p1(input)
	}
}

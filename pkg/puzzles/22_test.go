package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 22.
func TestDayTwentyTwoPartOne(t *testing.T) {
	input, err := file.ReadTestFile(22)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 37327623
	if r := d22p1(input); e != r {
		t.Errorf("d22p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 22.
func BenchmarkDayTwentyTwoPartOne(b *testing.B) {
	input, err := file.ReadInfile(22)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d22p1(input)
	}
}

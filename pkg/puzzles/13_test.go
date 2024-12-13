package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 13.
func TestDayThirteenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(13)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 480
	if r := d13p1(input); e != r {
		t.Errorf("d13p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 13.
func BenchmarkDayThirteenPartOne(b *testing.B) {
	input, err := file.ReadInfile(13)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d13p1(input)
	}
}

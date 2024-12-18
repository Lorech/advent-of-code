package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 18.
func TestDayEighteenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(18)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 22
	if r := d18p1(input, 7, 7, 12); e != r {
		t.Errorf("d18p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 18.
func BenchmarkDayEighteenPartOne(b *testing.B) {
	input, err := file.ReadInfile(18)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d18p1(input)
	}
}

package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 20.
func TestDayTwentyPartOne(t *testing.T) {
	input, err := file.ReadTestFile(20)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 44
	if r := d20p1(input, 1); e != r {
		t.Errorf("d20p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 20.
func BenchmarkDayTwentyPartOne(b *testing.B) {
	input, err := file.ReadInfile(20)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d20p1(input)
	}
}

// Tests the second part of the puzzle for day 20.
func TestDayTwentyPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(20)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 285
	if r := d20p2(input, 50); e != r {
		t.Errorf("d20p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 20.
func BenchmarkDayTwentyPartTwo(b *testing.B) {
	input, err := file.ReadInfile(20)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d20p2(input)
	}
}

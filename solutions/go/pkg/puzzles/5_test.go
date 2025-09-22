package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 5.
func TestDayFivePartOne(t *testing.T) {
	input, err := file.ReadTestFile(5)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 143
	if r := d5p1(input); e != r {
		t.Errorf("d5p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 5.
func TestDayFivePartTwo(t *testing.T) {
	input, err := file.ReadTestFile(5)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 123
	if r := d5p2(input); e != r {
		t.Errorf("d5p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 5.
func BenchmarkDayFivePartOne(b *testing.B) {
	input, err := file.ReadInfile(5)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d5p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 5.
func BenchmarkDayFivePartTwo(b *testing.B) {
	input, err := file.ReadInfile(5)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d5p2(input)
	}
}

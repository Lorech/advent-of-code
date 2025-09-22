package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 3.
func TestDayThreePartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 3)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 161
	if r := d3p1(input); e != r {
		t.Errorf("d3p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 3.
func TestDayThreePartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2024, 3)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 48
	if r := d3p2(input); e != r {
		t.Errorf("d3p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 3.
func BenchmarkDayThreePartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 3)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d3p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 3.
func BenchmarkDayThreePartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 3)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d3p2(input)
	}
}

package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 12.
func TestDayTwelvePartOne(t *testing.T) {
	input, err := file.ReadTestFile(12)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 1930
	if r := d12p1(input); e != r {
		t.Errorf("d12p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 12.
func BenchmarkDayTwelvePartOne(b *testing.B) {
	input, err := file.ReadInfile(12)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d12p1(input)
	}
}

// Tests the second part of the puzzle for day 12.
func TestDayTwelvePartTwo(t *testing.T) {
	input, err := file.ReadTestFile(12)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 1206
	if r := d12p2(input); e != r {
		t.Errorf("d12p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 12.
func BenchmarkDayTwelvePartTwo(b *testing.B) {
	input, err := file.ReadInfile(12)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d12p2(input)
	}
}

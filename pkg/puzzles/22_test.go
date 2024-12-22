package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 22.
func TestDayTwentyTwoPartOne(t *testing.T) {
	input, err := file.ReadTestFile(22, "part1")

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

// Tests the second part of the puzzle for day 22.
func TestDayTwentyTwoPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(22, "part2")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 23
	if r := d22p2(input); e != r {
		t.Errorf("d22p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 22.
func BenchmarkDayTwentyTwoPartTwo(b *testing.B) {
	input, err := file.ReadInfile(22)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d22p2(input)
	}
}

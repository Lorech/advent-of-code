package puzzles

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 17.
func TestDaySeventeenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 17, "part1")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := "4,6,3,5,6,3,5,2,1,0"
	if r := d17p1(input); e != r {
		t.Errorf("d17p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 17.
func BenchmarkDaySeventeenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 17)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d17p1(input)
	}
}

// Tests the second part of the puzzle for day 17.
func TestDaySeventeenPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2024, 17, "part2")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 117440
	if r := d17p2(input); e != r {
		t.Errorf("d17p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 17.
func BenchmarkDaySeventeenPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 17)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d17p2(input)
	}
}

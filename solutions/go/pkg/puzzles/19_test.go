package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 19.
func TestDayNineteenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 19)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 6
	if r := d19p1(input); e != r {
		t.Errorf("d19p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 19.
func BenchmarkDayNineteenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 19)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d19p1(input)
	}
}

// Tests the second part of the puzzle for day 19.
func TestDayNineteenPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2024, 19)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 16
	if r := d19p2(input); e != r {
		t.Errorf("d19p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 19.
func BenchmarkDayNineteenPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 19)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d19p2(input)
	}
}

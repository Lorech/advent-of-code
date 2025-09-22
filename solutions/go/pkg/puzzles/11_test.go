package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 11.
func TestDayElevenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2024, 11)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 55312
	if r := d11p1(input); e != r {
		t.Errorf("d11p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 11.
func BenchmarkDayElevenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2024, 11)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d11p1(input)
	}
}

// Benchmarks the second part of the puzzle for day 11.
func BenchmarkDayElevenPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2024, 11)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d11p2(input)
	}
}

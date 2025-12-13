package aoc2025

import (
	"lorech/advent-of-code/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 11.
func TestDayElevenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(2025, 11, "part1")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 5
	if r := d11p1(input); e != r {
		t.Errorf("d11p1() = %v, expected %v", r, e)
	}
}

// Benchmarks the first part of the puzzle for day 11.
func BenchmarkDayElevenPartOne(b *testing.B) {
	input, err := file.ReadInfile(2025, 11)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d11p1(input)
	}
}

// Tests the second part of the puzzle for day 11.
func TestDayElevenPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(2025, 11, "part2")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 2
	if r := d11p2(input); e != r {
		t.Errorf("d11p2() = %v, expected %v", r, e)
	}
}

// Benchmarks the second part of the puzzle for day 11.
func BenchmarkDayElevenPartTwo(b *testing.B) {
	input, err := file.ReadInfile(2025, 11)

	if err != nil {
		b.Errorf("Could not read file: %v", err)
		return
	}

	b.ResetTimer()
	for range b.N {
		d11p2(input)
	}
}

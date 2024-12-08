package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 3.
func TestDayThreePartOne(t *testing.T) {
	input, err := file.ReadTestFile(3)

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
	input, err := file.ReadTestFile(3)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 48
	if r := d3p2(input); e != r {
		t.Errorf("d3p2() = %v, expected %v", r, e)
	}
}

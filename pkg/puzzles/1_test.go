package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 1.
func TestDayOnePartOne(t *testing.T) {
	input, err := file.ReadTestFile(1)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 11
	if r := d1p1(input); e != r {
		t.Errorf("d1p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 1.
func TestDayOnePartTwo(t *testing.T) {
	input, err := file.ReadTestFile(1)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 31
	if r := d1p2(input); e != r {
		t.Errorf("d1p2() = %v, expected %v", r, e)
	}
}

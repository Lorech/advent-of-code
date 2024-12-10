package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 10.
func TestDayTenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(10)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 36
	if r := d10p1(input); e != r {
		t.Errorf("d10p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 10.
func TestDayTenPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(10)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 81
	if r := d10p2(input); e != r {
		t.Errorf("d10p2() = %v, expected %v", r, e)
	}
}

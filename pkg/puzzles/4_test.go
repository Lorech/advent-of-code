package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 4.
func TestDayFourPartOne(t *testing.T) {
	input, err := file.ReadTestFile(4)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 18
	if r := d4p1(input); e != r {
		t.Errorf("d4p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 4.
func TestDayFourPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(4)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 9
	if r := d4p2(input); e != r {
		t.Errorf("d4p2() = %v, expected %v", r, e)
	}
}

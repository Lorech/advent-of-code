package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 6.
func TestDaySixPartOne(t *testing.T) {
	input, err := file.ReadTestFile(6)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 41
	if r := d6p1(input); e != r {
		t.Errorf("d6p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 6.
func TestDaySixPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(6)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 6
	if r := d6p2(input); e != r {
		t.Errorf("d6p2() = %v, expected %v", r, e)
	}
}

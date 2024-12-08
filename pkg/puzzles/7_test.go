package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 7.
func TestDaySevenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(7)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 3749
	if r := d7p1(input); e != r {
		t.Errorf("d7p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 7.
func TestDaySevenPartSecond(t *testing.T) {
	input, err := file.ReadTestFile(7)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 11387
	if r := d7p2(input); e != r {
		t.Errorf("d7p2() = %v, expected %v", r, e)
	}
}

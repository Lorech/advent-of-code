package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the short example of the puzzle for day 8.
func TestDayNinePartOneSample(t *testing.T) {
	input, err := file.ReadTestFile(9, "12345")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 60
	if r := d9p1(input); e != r {
		t.Errorf("d9p1() = %v, expected %v", r, e)
	}
}

// Tests the first part of the puzzle for day 9.
func TestDayNinePartOne(t *testing.T) {
	input, err := file.ReadTestFile(9)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 1928
	if r := d9p1(input); e != r {
		t.Errorf("d9p1() = %v, expected %v", r, e)
	}
}

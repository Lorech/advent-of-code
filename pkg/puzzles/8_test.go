package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first example of the puzzle for day 8 where two antennas are present.
func TestDayEightPartOneFirstSample(t *testing.T) {
	input, err := file.ReadTestFile(8, "aa")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 2
	if r := d8p1(input); e != r {
		t.Errorf("d8p1() = %v, expected %v", r, e)
	}
}

// Tests the second example of the puzzle for day 8 where three antennas are present.
func TestDayEightPartOneSecondSample(t *testing.T) {
	input, err := file.ReadTestFile(8, "aaa")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 4
	if r := d8p1(input); e != r {
		t.Errorf("d8p1() = %v, expected %v", r, e)
	}
}

// Tests the first part of the puzzle for day 8.
func TestDayEightPartOne(t *testing.T) {
	input, err := file.ReadTestFile(8)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 14
	if r := d8p1(input); e != r {
		t.Errorf("d8p1() = %v, expected %v", r, e)
	}
}

// Tests the fourth example of the puzzle for day 8 where T-antennas are present.
func TestDayEightPartTwoFirstSample(t *testing.T) {
	input, err := file.ReadTestFile(8, "TTT")

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 9
	if r := d8p2(input); e != r {
		t.Errorf("d8p2() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 8.
func TestDayEightPartTwo(t *testing.T) {
	input, err := file.ReadTestFile(8)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 34
	if r := d8p2(input); e != r {
		t.Errorf("d8p2() = %v, expected %v", r, e)
	}
}

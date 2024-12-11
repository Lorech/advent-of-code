package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the first part of the puzzle for day 11.
func TestDayElevenPartOne(t *testing.T) {
	input, err := file.ReadTestFile(11)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	e := 55312
	if r := d11p1(input); e != r {
		t.Errorf("d11p1() = %v, expected %v", r, e)
	}
}

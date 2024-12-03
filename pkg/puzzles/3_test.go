package puzzles

import "testing"

var inputDayThree string = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

// Tests the first part of the puzzle for day 3.
func TestDayThreePartOne(t *testing.T) {
	e := 161
	if r := d3p1(inputDayThree); e != r {
		t.Errorf("d3p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 3.
func TestDayThreePartTwo(t *testing.T) {
	e := 48
	if r := d3p2(inputDayThree); e != r {
		t.Errorf("d3p2() = %v, expected %v", r, e)
	}
}

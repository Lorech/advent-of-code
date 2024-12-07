package puzzles

import "testing"

var inputDaySeven string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

// Tests the first part of the puzzle for day 7.
func TestDaySevenPartOne(t *testing.T) {
	e := 3749
	if r := d7p1(inputDaySeven); e != r {
		t.Errorf("d7p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 7.
func TestDaySevenPartSecond(t *testing.T) {
	e := 11387
	if r := d7p2(inputDaySeven); e != r {
		t.Errorf("d7p2() = %v, expected %v", r, e)
	}
}

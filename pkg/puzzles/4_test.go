package puzzles

import "testing"

var inputDayFour string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

// Tests the first part of the puzzle for day 4.
func TestDayFourPartOne(t *testing.T) {
	e := 18
	if r := d4p1(inputDayFour); e != r {
		t.Errorf("d4p1() = %v, expected %v", r, e)
	}
}

// Tests the secodn part of the puzzle for day 4.
func TestDayFourPartTwo(t *testing.T) {
	e := 9
	if r := d4p2(inputDayFour); e != r {
		t.Errorf("d4p2() = %v, expected %v", r, e)
	}
}

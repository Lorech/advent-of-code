package puzzles

import "testing"

var inputDaySix string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

// Tests the first part of the puzzle for day 6.
func TestDaySixPartOne(t *testing.T) {
	e := 41
	if r := d6p1(inputDaySix); e != r {
		t.Errorf("d6p1() = %v, expected %v", r, e)
	}
}

// Tests the second part of the puzzle for day 6.
func TestDaySixPartTwo(t *testing.T) {
	e := 6
	if r := d6p2(inputDaySix); e != r {
		t.Errorf("d6p2() = %v, expected %v", r, e)
	}
}

func TestDaySixEdgeCases(t *testing.T) {
	e := 0
	if r := d6p2(`##..
...#
....
^.#.`); e != r {
		t.Errorf("d6p1() = %v, expected %v", r, e)
	}
}

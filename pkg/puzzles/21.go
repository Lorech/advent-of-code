package puzzles

import (
	"fmt"
	"lorech/advent-of-code-2024/pkg/cslices"
	"lorech/advent-of-code-2024/pkg/grid"
	"slices"
	"strconv"
	"strings"
	"time"
)

type keypad struct {
	keys     map[rune]grid.Coordinates // A mapping of keypad characters to their coordinates in the grid.
	position grid.Coordinates          // The position within the grid that is currently active.
	focus    rune                      // The currently focused character on the keypad based on position.
	cache    map[[2]rune]int           // Memoized storage for number of moves to get from the first rune to the second rune.
}

// Initializes a new instance of a numpad-like keypad.
func newNumpad() keypad {
	var k keypad
	k.cache = make(map[[2]rune]int, 0)
	k.keys = map[rune]grid.Coordinates{
		'7': grid.Coordinates{X: 0, Y: 0}, '8': grid.Coordinates{X: 1, Y: 0}, '9': grid.Coordinates{X: 2, Y: 0},
		'4': grid.Coordinates{X: 0, Y: 1}, '5': grid.Coordinates{X: 1, Y: 1}, '6': grid.Coordinates{X: 2, Y: 1},
		'1': grid.Coordinates{X: 0, Y: 2}, '2': grid.Coordinates{X: 1, Y: 2}, '3': grid.Coordinates{X: 2, Y: 2},
		' ': grid.Coordinates{X: 0, Y: 3}, '0': grid.Coordinates{X: 1, Y: 3}, 'A': grid.Coordinates{X: 2, Y: 3},
	}
	k.position = grid.Coordinates{X: 2, Y: 3}
	k.focus = 'A'
	return k
}

// Initializes a new instance of an arrow key style keypad.
func newDirectionPad() keypad {
	var k keypad
	k.cache = make(map[[2]rune]int, 0)
	k.keys = map[rune]grid.Coordinates{
		' ': grid.Coordinates{X: 0, Y: 0}, '^': grid.Coordinates{X: 1, Y: 0}, 'A': grid.Coordinates{X: 2, Y: 0},
		'<': grid.Coordinates{X: 0, Y: 1}, 'v': grid.Coordinates{X: 1, Y: 1}, '>': grid.Coordinates{X: 2, Y: 1},
	}
	k.position = grid.Coordinates{X: 2, Y: 0}
	k.focus = 'A'
	return k
}

// Moves the keypad to a different position, returning a slice of moves to get there.
func (k *keypad) move(char rune) []rune {
	coordinates, found := k.keys[char]

	if !found {
		panic("Wrong keypad!")
	}

	moves := make([]rune, 0)
	xd, yd := coordinates.X-k.position.X, coordinates.Y-k.position.Y
	horizontal, vertical := '<', '^'
	if xd > 0 {
		horizontal = '>'
	}
	if yd > 0 {
		vertical = 'v'
	}

	// We want to maximize sequential moves, and we want to prioritize moves that
	// are as far away from the starting point as possible, as that helps the first
	// point execute in higher order chains, so:
	//   - Only move in one direction at a time as zigzagging would create large
	// amounts of moves in higher tiers;
	//   - Prefer moving left, then move vertically, then move right, as that
	// creates the largest chain of matching moves in higher tiers.
	if xd < 0 && k.keys[' '].Y == k.position.Y && k.keys[' '].X == k.position.X+xd || xd >= 0 {
		moves = append(moves, slices.Repeat([]rune{vertical}, max(yd, -yd))...)
		moves = append(moves, slices.Repeat([]rune{horizontal}, max(xd, -xd))...)
	} else {
		moves = append(moves, slices.Repeat([]rune{horizontal}, max(xd, -xd))...)
		moves = append(moves, slices.Repeat([]rune{vertical}, max(yd, -yd))...)
	}
	moves = append(moves, 'A')

	k.position = coordinates
	k.focus = char

	return moves
}

// Day 21: Keypad Conundrum
// https://adventofcode.com/2024/day/21
func dayTwentyOne(input string) (int, int) {
	return d21p1(input), 0
}

// Completes the first half of the puzzle for day 21.
func d21p1(input string) int {
	codes := strings.Split(input, "\n")
	numpad := newNumpad()
	controls := make([]keypad, 3)
	for i := range controls {
		controls[i] = newDirectionPad()
	}

	complexity := 0
	for _, code := range codes {
		start := time.Now()
		c := calculateCodeComplexity(code, 3, numpad, &controls)
		end := time.Now()
		fmt.Printf("Code %s has complexity %d, calculated in %v\n", code, c, end.Sub(start))
		complexity += c
	}
	return complexity
}

// Calculates the code complexity for a single code at the provided robot depth.
func calculateCodeComplexity(code string, maxDepth int, numpad keypad, controls *[]keypad) int {
	moves := enterCode(code, maxDepth, numpad, controls)
	nums := make([]int, len(code)-1)
	for i, num := range code[:len(code)-1] {
		n, _ := strconv.Atoi(string(num))
		nums[i] = n
	}
	value, _ := cslices.Stoi(nums)
	return value * moves
}

// Enters the provided keycode using layered seperation up to maxDepth away,
// returning the number of moves required to enter the code.
func enterCode(code string, maxDepth int, numpad keypad, controls *[]keypad) int {
	n := 0

	for _, char := range code {
		moves := numpad.move(char)

		// FIXME: Ugly, but this is the way I want this functions' API to work.
		if maxDepth > 0 {
			for _, move := range moves {
				navigateNumpad(&n, move, 1, maxDepth, controls)
			}
		} else {
			return len(moves)
		}
	}

	return n
}

// Recursively navigates a numpad through several layers of directional inputs,
// storing the final number of moves to produce the required character in `n`.
func navigateNumpad(n *int, button rune, depth int, maxDepth int, controls *[]keypad) {
	if depth == maxDepth {
		*n++
		return
	}

	cacheKey := [2]rune{(*controls)[depth].focus, button}
	memo, cached := (*controls)[depth].cache[cacheKey]
	if cached {
		*n += memo
		(*controls)[depth].focus = button
		(*controls)[depth].position = (*controls)[depth].keys[button]

		// Move all other controls to the starting position, as they must have hit A to trigger this move.
		for i := depth + 1; i < maxDepth; i++ {
			(*controls)[i].focus = 'A'
			(*controls)[i].position = (*controls)[i].keys['A']
		}
		return
	}

	moves := (*controls)[depth].move(button)
	nested := 0
	for _, move := range moves {
		navigateNumpad(&nested, move, depth+1, maxDepth, controls)
	}

	(*controls)[depth].cache[cacheKey] = nested
	*n += nested
}

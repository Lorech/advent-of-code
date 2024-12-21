package puzzles

import (
	"lorech/advent-of-code-2024/pkg/cslices"
	"lorech/advent-of-code-2024/pkg/grid"
	"slices"
	"strconv"
	"strings"
)

type keypad struct {
	keys     map[rune]grid.Coordinates // A mapping of keypad characters to their coordinates in the grid.
	position grid.Coordinates          // The position within the grid that is currently active.
}

// Moves the keypad to a different position, returning a slice of moves to get there.
func (k *keypad) move(char rune) []rune {
	coordinates, found := k.keys[char]

	if !found {
		panic("Wrong keypad!")
	}

	moves := make([]rune, 0)
	xd, yd := coordinates.X-k.position.X, coordinates.Y-k.position.Y
	horizontal, vertical := '<', 'v' // Prioritize farther away moves by default.
	if xd > 0 {
		horizontal = '>'
	}
	if yd < 0 {
		vertical = '^'
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
	return moves
}

// Day 21: Keypad Conundrum
// https://adventofcode.com/2024/day/21
func dayTwentyOne(input string) (int, int) {
	return d21p1(input), 0
}

// Completes the first half of the puzzle for day 21.
func d21p1(input string) int {
	codes, numpad, controls := parseCodes(input)
	complexity := 0

	for _, code := range codes {
		moves := enterCode(code, 3, numpad, controls)
		nums := make([]int, len(code)-1)
		for i, num := range code[:len(code)-1] {
			n, _ := strconv.Atoi(string(num))
			nums[i] = n
		}
		value, _ := cslices.Stoi(nums)
		complexity += value * len(moves)
	}

	return complexity
}

// Enters the provided keycode using layered seperation up to maxDepth away.
func enterCode(code []rune, maxDepth int, numpad keypad, controls keypad) []rune {
	inputs := make([]rune, 0)
	controlPads := slices.Repeat([]keypad{controls}, maxDepth)

	for _, char := range code {
		moves := numpad.move(char)

		// FIXME: Ugly, but this is the way I want this function's API to work.
		if maxDepth > 0 {
			for _, move := range moves {
				navigateNumpad(&inputs, move, 1, maxDepth, &controlPads)
			}
		} else {
			return moves
		}
	}

	return inputs
}

// Recursively navigates a numpad through several layers of directional inputs,
// storing the final list of moves to produce the required character in `inputs`.
func navigateNumpad(inputs *[]rune, button rune, depth int, maxDepth int, controls *[]keypad) {
	if depth == maxDepth {
		*inputs = append(*inputs, button)
		return
	}

	moves := (*controls)[depth].move(button)
	for _, move := range moves {
		navigateNumpad(inputs, move, depth+1, maxDepth, controls)
	}
}

// Parses the input data into individual characters of each keycode to enter,
// and centralizes the initialization of the keypads to the root of the solution.
func parseCodes(input string) ([][]rune, keypad, keypad) {
	rows := strings.Split(input, "\n")
	codes := make([][]rune, len(rows))

	for i, row := range rows {
		codes[i] = make([]rune, len(row))
		for j, char := range row {
			codes[i][j] = char
		}
	}

	numpad := keypad{
		keys: map[rune]grid.Coordinates{
			'7': grid.Coordinates{X: 0, Y: 0}, '8': grid.Coordinates{X: 1, Y: 0}, '9': grid.Coordinates{X: 2, Y: 0},
			'4': grid.Coordinates{X: 0, Y: 1}, '5': grid.Coordinates{X: 1, Y: 1}, '6': grid.Coordinates{X: 2, Y: 1},
			'1': grid.Coordinates{X: 0, Y: 2}, '2': grid.Coordinates{X: 1, Y: 2}, '3': grid.Coordinates{X: 2, Y: 2},
			' ': grid.Coordinates{X: 0, Y: 3}, '0': grid.Coordinates{X: 1, Y: 3}, 'A': grid.Coordinates{X: 2, Y: 3},
		},
		position: grid.Coordinates{X: 2, Y: 3},
	}

	controls := keypad{
		keys: map[rune]grid.Coordinates{
			' ': grid.Coordinates{X: 0, Y: 0}, '^': grid.Coordinates{X: 1, Y: 0}, 'A': grid.Coordinates{X: 2, Y: 0},
			'<': grid.Coordinates{X: 0, Y: 1}, 'v': grid.Coordinates{X: 1, Y: 1}, '>': grid.Coordinates{X: 2, Y: 1},
		},
		position: grid.Coordinates{X: 2, Y: 0},
	}

	return codes, numpad, controls
}

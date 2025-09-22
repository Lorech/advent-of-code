package puzzles

import (
	"lorech/advent-of-code-2024/pkg/grid"
	"slices"
	"strings"
)

// Day 15: Warehouse Woes
// https://adventofcode.com/2024/day/15
func dayFifteen(input string) (int, int) {
	return d15p1(input), d15p2(input)
}

// Completes the first half of the puzzle for day 15.
func d15p1(input string) int {
	warehouse, guard, moves := parseWarehouse(input)

	// Move the robot through his routine.
	for _, move := range moves {
		yd, xd := move.Velocity()
		objects := 0
		y, x := guard[0]+yd, guard[1]+xd
		valid, moved := false, false

		for moved == false {
			switch warehouse[y][x] {
			case '#':
				// This sends us into a wall, so nothing happens.
				moved = true
			case 'O':
				// We would push something, which only throws off the step counter.
				objects++
				y += yd
				x += xd
				continue
			case '.':
				// Free space! We can make a move, so exit and do it.
				valid = true
				moved = true
			}
		}

		if valid {
			// Move the guard forward 1 step.
			warehouse[guard[0]][guard[1]] = '.'
			warehouse[guard[0]+yd][guard[1]+xd] = '@'
			guard[0] += yd
			guard[1] += xd

			// If we were pushing anything, move the last object forward 1 step.
			// If we weren't pushing an object, the guard is on this tile, so do nothing.
			if objects != 0 {
				warehouse[y][x] = 'O'
			}
		}
	}

	// Calculate the position of each box.
	coordinates := 0
	for y, row := range warehouse {
		start := 1
		for true {
			ob := slices.Index(row[start:], 'O')
			if ob == -1 {
				break
			}

			coordinates += 100*y + start + ob
			start += ob + 1
		}
	}

	return coordinates
}

// Completes the second half of the puzzle for day 15.
func d15p2(input string) int {
	warehouse, guard, moves := parseWarehouse(input, true)

	// Move the robot through his routine.
	for _, move := range moves {
		yd, xd := move.Velocity()
		stack := make([]grid.Tile, 1)
		stack[0] = grid.Tile{
			Position: grid.Coordinates{
				X: guard[1],
				Y: guard[0],
			},
			Value: '@',
		}
		i := 0

		// Check if there is a valid move to be made for every element in the stack.
		// Increment instead of going through the stack as that wouldn't update the loop condition.
		for i < len(stack) {
			t := stack[i]
			i++

			// Don't parse the inside half of an object when moving horizontally
			// to prevent adding false objects in the stack.
			if t.Value == '[' && xd > 0 || t.Value == ']' && xd < 0 {
				continue
			}

			y, x := t.Position.Y+yd, t.Position.X+xd
			switch warehouse[y][x] {
			case '#':
				// This tile would go into the wall, so do nothing this turn.
				goto nextMove
			case '.':
				// The current tile can be moved, but we need to check if the others can.
				continue
			case '[':
				// Left half of an object - we will need to check if it can be moved.
				stack = append(stack, grid.Tile{
					Position: grid.Coordinates{
						X: x,
						Y: y,
					},
					Value: '[',
				})
				stack = append(stack, grid.Tile{
					Position: grid.Coordinates{
						X: x + 1,
						Y: y,
					},
					Value: ']',
				})
			case ']':
				// Right half of an object - we will need to check if it can be moved.
				stack = append(stack, grid.Tile{
					Position: grid.Coordinates{
						X: x,
						Y: y,
					},
					Value: ']',
				})
				stack = append(stack, grid.Tile{
					Position: grid.Coordinates{
						X: x - 1,
						Y: y,
					},
					Value: '[',
				})
			}
		}

		// If we made it out of the for-loop, we can move all of the tiles!
		// Do it backwards, because we need to do LIFO to avoid overwriting tiles.
		for i := len(stack) - 1; i >= 0; i-- {
			t := stack[i]
			warehouse[t.Position.Y+yd][t.Position.X+xd] = t.Value
			warehouse[t.Position.Y][t.Position.X] = '.'
		}
		guard[0] += yd
		guard[1] += xd
	nextMove: // Break out of the loop while avoiding value updates in failure cases.
	}

	// Calculate the position of each box.
	coordinates := 0
	for y, row := range warehouse {
		start := 1
		for true {
			ob := slices.Index(row[start:], '[')
			if ob == -1 {
				break
			}

			coordinates += 100*y + start + ob
			start += ob + 1
		}
	}

	return coordinates
}

// Parses the input data into a grid representing the warehouse's initial
// state, the robot's starting position, and a slice containing all the moves
// the robot is supposed to make.
func parseWarehouse(input string, scaled ...bool) ([][]rune, [2]int, []grid.Direction) {
	rows := strings.Split(input, "\n")
	warehouse := make([][]rune, 0)
	moves := make([]grid.Direction, 0)
	var robot [2]int

	for y, row := range rows {
		// Delimiting row between warehouse and moves.
		if len(row) == 0 {
			continue
		}

		// Warehouse grid.
		if row[0] == '#' {
			// Scale the warehouse by a factor of 2 if it's required by the puzzle.
			if len(scaled) > 0 && scaled[0] {
				r := make([]rune, len(row)*2)
				for i, tile := range row {
					x := i * 2
					switch tile {
					case '#':
						r[x] = '#'
						r[x+1] = '#'
					case 'O':
						r[x] = '['
						r[x+1] = ']'
					case '.':
						r[x] = '.'
						r[x+1] = '.'
					case '@':
						r[x] = '@'
						r[x+1] = '.'
						robot = [2]int{y, x}
					}
				}
				warehouse = append(warehouse, []rune(r))

				// Since we parsed each tile manually, we are done processing the row.
				continue
			}

			// Check if this row contains the robot's coordinates if we dont know
			// already know them. We can check for 0 because the warehouse has a
			// border, meaning he can never spawn at 0,0 to start off.
			if robot[0] == 0 {
				x := strings.IndexRune(row, '@')
				if x != -1 {
					robot = [2]int{y, x}
				}
			}

			warehouse = append(warehouse, []rune(row))
			continue
		}

		// Move row.
		m := make([]grid.Direction, len(row))
		for x, move := range row {
			switch move {
			case '^':
				m[x] = grid.Up
			case 'v':
				m[x] = grid.Down
			case '<':
				m[x] = grid.Left
			case '>':
				m[x] = grid.Right
			}
		}
		moves = append(moves, m...)
	}

	return warehouse, robot, moves
}

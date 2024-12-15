package puzzles

import (
	"lorech/advent-of-code-2024/pkg/grid"
	"slices"
	"strings"
)

// Day 15: Warehouse Woes
// https://adventofcode.com/2024/day/15
func dayFifteen(input string) (int, int) {
	return d15p1(input), 0
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

// Parses the input data into a grid representing the warehouse's initial
// state, the robot's starting position, and a slice containing all the moves
// the robot is supposed to make.
func parseWarehouse(input string) ([][]rune, [2]int, []grid.Direction) {
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

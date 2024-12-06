package puzzles

import (
	"regexp"
	"slices"
	"strings"
)

type pointingDirection int

const (
	Up pointingDirection = iota
	Down
	Left
	Right
)

var pointing = map[pointingDirection]string{
	Up:    "up",
	Down:  "down",
	Left:  "left",
	Right: "right",
}

func rotate(p pointingDirection) pointingDirection {
	switch p {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}

	panic("Invalid pointing direction!")
}

func moving(p pointingDirection) (int, int) {
	var (
		yd, xd int
	)

	switch p {
	case Up:
		yd = -1
		xd = 0
	case Down:
		yd = 1
		xd = 0
	case Left:
		yd = 0
		xd = -1
	case Right:
		yd = 0
		xd = 1
	}

	return yd, xd
}

// Day 6: Guard Gallivant
// https://adventofcode.com/2024/day/6
func daySix(input string) (int, int) {
	return d6p1(input), 0
}

// Completes the first half of the puzzle for day 6.
func d6p1(input string) int {
	room, guard := parse(input)
	pointing := Up
	visited := make(map[int][]int)

	// NOTE: Intentionally doing it the crappy way in case it comes in handy for part 2.
	for true {
		yd, xd := moving(pointing)

		for steps := 0; true; steps++ {
			x := guard[1] + xd*steps
			y := guard[0] + yd*steps

			// If this is a new tile, add it to the mapping.
			_, exists := visited[y]
			if exists {
				if !slices.Contains(visited[y], x) {
					visited[y] = append(visited[y], x)
				}
			} else {
				visited[y] = []int{x}
			}

			// Check if we're about to go out of the room.
			if y+yd < 0 || y+yd >= len(room) || x+xd < 0 || x+xd >= len(room[1]) {
				goto leaveRoom
			}

			// If we run into a dead-end, rotate and restart our steps counter.
			if room[y+yd][x+xd] == '#' {
				pointing = rotate(pointing)
				guard[0] = y
				guard[1] = x
				break
			}
		}
	}

leaveRoom:
	totalVisits := 0
	for _, tiles := range visited {
		totalVisits += len(tiles)
	}
	return totalVisits
}

// Parses the input data, returning:
// - A 2D slice of the room's layout in y,x space, replacing the guard with a period;
// - The guard's position within the room in y,x space.
func parseLaboratory(input string) ([][]byte, [2]int) {
	// Find the position of the guard within the data and remove him from the room.
	gi := strings.Index(input, "^")
	input = strings.Replace(input, "^", ".", -1)

	// Parse the room into a 2D slice of y,x coordinates.
	reRows := regexp.MustCompile(`[.#]+`)
	rows := reRows.FindAll([]byte(input), -1)
	room := make([][]byte, len(rows))
	for i, row := range rows {
		room[i] = make([]byte, len(row))
		room[i] = row
	}

	// Find the coordinates of the guard within the room.
	gy := gi / (len(rows[0]) + 1)
	gx := gi - (gy)*(len(rows[0])+1)
	guard := [2]int{gy, gx}

	return room, guard
}

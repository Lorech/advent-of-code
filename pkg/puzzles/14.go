package puzzles

import (
	"regexp"
	"strconv"
	"strings"
)

type guard struct {
	P coordinates // The position of the guard.
	V coordinates // The velocity of the guard.
}

type coordinates struct {
	X int
	Y int
}

// Day 14: Restroom Redoubt
// https://adventofcode.com/2024/day/14
func dayFourteen(input string) (int, int) {
	return d14p1(input), 0
}

// Completes the first half of the puzzle for day 14.
func d14p1(input string, roomSize ...int) int {
	guards := parseGuards(input)
	maxX := 101
	maxY := 103

	if len(roomSize) > 0 {
		maxX = roomSize[0]
		maxY = roomSize[1]
	}

	middleX := (maxX - 1) / 2
	middleY := (maxY - 1) / 2

	// Loop for 100 "seconds" to determine where the guards end up at.
	for range 100 {
		passSecond(&guards, maxX, maxY)
	}

	// Determine the number of guards in each quadrant.
	quadrants := [4][]guard{}
	for _, guard := range guards {
		if guard.P.X < middleX && guard.P.Y < middleY {
			quadrants[0] = append(quadrants[0], guard)
		} else if guard.P.X < middleX && guard.P.Y > middleY {
			quadrants[1] = append(quadrants[1], guard)
		} else if guard.P.X > middleX && guard.P.Y < middleY {
			quadrants[2] = append(quadrants[2], guard)
		} else if guard.P.X > middleX && guard.P.Y > middleY {
			quadrants[3] = append(quadrants[3], guard)
		} else {
			// Do nothing if the guard does not fit a specific quadrant.
		}
	}

	return len(quadrants[0]) * len(quadrants[1]) * len(quadrants[2]) * len(quadrants[3])
}

// Moves every guard for 1 second worth of momvement.
func passSecond(guards *[]guard, maxX int, maxY int) {
	for j, guard := range *guards {
		guard.P.X += guard.V.X
		if guard.P.X >= maxX {
			guard.P.X -= maxX
		} else if guard.P.X < 0 {
			guard.P.X += maxX
		}

		guard.P.Y += guard.V.Y
		if guard.P.Y >= maxY {
			guard.P.Y -= maxY
		} else if guard.P.Y < 0 {
			guard.P.Y += maxY
		}

		(*guards)[j] = guard
	}
}

// Parses the input data into a slice of guards in the bathroom.
func parseGuards(input string) []guard {
	rows := strings.Split(input, "\n")
	reData := regexp.MustCompile(`(-?\d+),(-?\d+)`)
	guards := make([]guard, 0)

	for _, row := range rows {
		coords := reData.FindAllString(row, -1)
		p := strings.Split(coords[0], ",")
		v := strings.Split(coords[1], ",")
		px, _ := strconv.Atoi(p[0])
		py, _ := strconv.Atoi(p[1])
		vx, _ := strconv.Atoi(v[0])
		vy, _ := strconv.Atoi(v[1])
		guards = append(guards, guard{P: coordinates{px, py}, V: coordinates{vx, vy}})
	}

	return guards
}

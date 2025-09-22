package puzzles

import (
	"lorech/advent-of-code/pkg/cmath"
	"math"
	"regexp"
	"slices"
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
	return d14p1(input), d14p2(input)
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

// Completes the second half of the puzzle for day 14.
func d14p2(input string) int {
	guards := parseGuards(input)
	seconds := 0
	xVariances, yVariances := make([]float64, 0), make([]float64, 0)

	for true {
		// Find the sum of both coordinates for every guard.
		var (
			xSum, ySum float64
		)
		for _, guard := range guards {
			xSum += float64(guard.P.X)
			ySum += float64(guard.P.Y)
		}

		// Find the mean of both coordinates.
		xAvg, yAvg := xSum/float64(len(guards)), ySum/float64(len(guards))
		var (
			xDiffs, yDiffs float64
		)
		for _, guard := range guards {
			xDiffs += math.Pow(float64(guard.P.X)-xAvg, 2)
			yDiffs += math.Pow(float64(guard.P.Y)-yAvg, 2)
		}

		// Calculate the variance in the guards' positions.
		xVariance, yVariance := xDiffs/float64(len(guards)), yDiffs/float64(len(guards))

		// If we have no previous variances, this would be an outlier by default.
		if len(xVariances) < 2 {
			xVariances = append(xVariances, xVariance)
			yVariances = append(yVariances, yVariance)
			seconds++
			passSecond(&guards, 101, 103)
			continue
		}

		// Find the IQR of the known variances.
		slices.Sort(xVariances)
		slices.Sort(yVariances)
		xMid, yMid := len(xVariances)/2, len(yVariances)/2
		xOffset, yOffset := 0, 0
		if len(xVariances)%2 == 1 {
			xOffset++
		}
		if len(yVariances)%2 == 1 {
			yOffset++
		}
		xQ1 := cmath.Median(xVariances[:xMid])
		xQ3 := cmath.Median(xVariances[xMid+xOffset:])
		yQ1 := cmath.Median(yVariances[:yMid])
		yQ3 := cmath.Median(yVariances[yMid+yOffset:])

		// Check if the variance in these values is an outlier between previous iterations.
		xIQR := xQ3 - xQ1
		yIQR := yQ3 - yQ1
		if (xVariance < xQ1-1.5*xIQR || xVariance > xQ3+1.5*xIQR) && (yVariance < yQ1-1.5*yIQR || yVariance > yQ3+1.5*yIQR) {
			// This is an outlier! We found the tree!
			break
		}

		// This is not an outlier, so add it for the next iteration.
		xVariances = append(xVariances, xVariance)
		yVariances = append(yVariances, yVariance)
		seconds++
		passSecond(&guards, 101, 103)
	}

	return seconds
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

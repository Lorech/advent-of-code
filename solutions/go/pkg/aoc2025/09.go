package aoc2025

import (
	"lorech/advent-of-code/pkg/grid"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Day 9: Movie Theater
// https://adventofcode.com/2025/day/9
func dayNine(input string) (int, int) {
	return d9p1(input), d9p2(input)
}

// Completes the first half of the puzzle for day 9.
func d9p1(input string) int {
	tiles, _, _ := parseTiles(input)
	maxArea := 0

	for i := 0; i < len(tiles)-1; i++ {
		a := tiles[i]
		for j := i + 1; j < len(tiles); j++ {
			b := tiles[j]
			area := rectangleArea(a, b)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

// Completes the second half of the puzzle for day 9.
func d9p2(input string) int {
	tiles, _, _ := parseTiles(input)
	maxArea := 0

	// Support looping through the tile set two tiles at a time to get every line
	tiles = append(tiles, tiles[0])

	outline := make([]grid.Coordinate, 0)
	for i := 0; i < len(tiles)-1; i += 2 {
		s, e := tiles[i], tiles[i+1]
		connectTiles(s, e, &outline)
	}

	for i := 0; i < len(tiles)-1; i++ {
		a := tiles[i]
		for j := i + 1; j < len(tiles); j++ {
			b := tiles[j]
			area := rectangleArea(a, b)
			if area > maxArea {
				c, d := grid.Coordinate{X: a.X, Y: b.Y}, grid.Coordinate{X: b.X, Y: a.Y}
				tl, tr, br, bl := alignRectangle(a, b, c, d)
				top, bot, lft, rgt, valid := []grid.Coordinate{tl, tr}, []grid.Coordinate{bl, br},
					[]grid.Coordinate{tl, bl}, []grid.Coordinate{tr, br}, true

				for k := 0; k < len(tiles)-1; k++ {
					l1, l2 := tiles[k], tiles[k+1]
					vertical := l1.X == l2.X

					// If the static axis is on an edge or outside the rectangle, it's safe
					if vertical {
						if l1.X <= tl.X || l1.X >= tr.X {
							continue
						}
					} else {
						if l1.Y <= tl.Y || l1.Y >= bl.Y {
							continue
						}
					}

					var p1i, p2i bool
					if vertical {
						_, p1i = lineSegmentIntersection(top[0], top[1], l1, l2)
						_, p2i = lineSegmentIntersection(bot[0], bot[1], l1, l2)
					} else {
						_, p1i = lineSegmentIntersection(lft[0], lft[1], l1, l2)
						_, p2i = lineSegmentIntersection(rgt[0], rgt[1], l1, l2)
					}

					// If both edges are intersected, the rectangle is broken in half
					if p1i && p2i {
						valid = false
						break
					}

					// If neither edge intersects, this line is not in contact with the rectangle
					if !p1i && !p2i {
						continue
					}

					// If one edge intersects and one point is inside the rectangle,
					// the line creates a cutout in the rectangle, breaking it.
					if vertical {
						if (l1.Y > tl.Y && l1.Y < bl.Y) || (l2.Y > tl.Y && l2.Y < bl.Y) {
							valid = false
							break
						}
					} else {
						if (l1.X > tl.X && l1.X < tr.X) || (l2.X > tl.X && l2.X < tr.X) {
							valid = false
							break
						}
					}
				}

				if valid {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// Check if two line segments a and b, located on the same plane, intersect each other.
// Returns the intersection point and a flag if the intersection happened.
// TODO: Candidate for separating into a package function.
func lineSegmentIntersection(a1, a2, b1, b2 grid.Coordinate) (grid.Coordinate, bool) {
	p, i := grid.Coordinate{X: 0, Y: 0}, false

	d := (a1.X-a2.X)*(b1.Y-b2.Y) - (a1.Y-a2.Y)*(b1.X-b2.X)
	if d == 0 {
		return p, i
	}

	t := float64((a1.X-b1.X)*(b1.Y-b2.Y)-(a1.Y-b1.Y)*(b1.X-b2.X)) / float64(d)
	u := -float64((a1.X-a2.X)*(a1.Y-b1.Y)-(a1.Y-a2.Y)*(a1.X-b1.X)) / float64(d)

	i = 0 <= t && t <= 1 && 0 <= u && u <= 1
	if i {
		p.X = (a1.X + int(t*float64(a2.X-a1.X)))
		p.Y = (a1.Y + int(t*float64(a2.Y-a1.Y)))
	}

	return p, i
}

// Check if two lines a and b, located on the same plane, intersect each other.
// Returns the intersection point and a flag if the intersection happened.
// TODO: Candidate for separating into a package function.
func lineIntersection(a1, a2, b1, b2 grid.Coordinate) (grid.Coordinate, bool) {
	p, i := grid.Coordinate{X: 0, Y: 0}, false

	d := (a1.X-a2.X)*(b1.Y-b2.Y) - (a1.Y-a2.Y)*(b1.X-b2.X)
	if d != 0 {
		pxn := (a1.X*a2.Y-a1.Y*a2.X)*(b1.X-b2.X) - (a1.X-a2.X)*(b1.X*b2.Y-b1.Y*b2.X)
		pyn := (a1.X*a2.Y-a1.Y*a2.X)*(b1.Y-b2.Y) - (a1.Y-a2.Y)*(b1.X*b2.Y-b1.Y*b2.X)
		p.X = pxn / d
		p.Y = pyn / d
		i = true
	}

	return p, i
}

// Aligns four arbitrarily order corners of a rectangle into a predictable
// set of (in order) top left, top right, bottom right, bottom left corners.
// TODO: Candidate for separating into a package function.
func alignRectangle(a, b, c, d grid.Coordinate) (grid.Coordinate, grid.Coordinate, grid.Coordinate, grid.Coordinate) {
	xs, ys := []int{a.X, b.X, c.X, d.X}, []int{a.Y, b.Y, c.Y, d.Y}
	slices.Sort(xs)
	slices.Sort(ys)
	minX, minY, maxX, maxY := xs[0], ys[0], xs[3], ys[3]
	return grid.Coordinate{X: minX, Y: minY},
		grid.Coordinate{X: maxX, Y: minY},
		grid.Coordinate{X: maxX, Y: maxY},
		grid.Coordinate{X: minX, Y: maxY}
}

// Calculates the area of a rectangle based on its two opposing corners on a grid.
// TODO: Candidate for separating into a package function.
func rectangleArea(a, b grid.Coordinate) int {
	return int((math.Abs(float64(a.X-b.X)) + 1) * (math.Abs(float64(a.Y-b.Y)) + 1))
}

// Connects two tiles together with a straight line of tiles
func connectTiles(a, b grid.Coordinate, tiles *[]grid.Coordinate) {
	if a.X < b.X {
		for x := a.X; x <= b.X; x++ {
			*tiles = append(*tiles, grid.Coordinate{X: x, Y: a.Y})
		}
	} else {
		for x := a.X; x >= b.X; x-- {
			*tiles = append(*tiles, grid.Coordinate{X: x, Y: a.Y})
		}
	}
	if a.Y < b.Y {
		for y := a.Y; y <= b.Y; y++ {
			*tiles = append(*tiles, grid.Coordinate{X: a.X, Y: y})
		}
	} else {
		for y := a.Y; y >= b.Y; y-- {
			*tiles = append(*tiles, grid.Coordinate{X: a.X, Y: y})
		}
	}
}

// Parses input data into structured data.
// Returns a tuple containing:
// - a slice of every red tile within the grid
// - the maximum X position of a red tile within a 0-indexed grid
// - the maximum Y position of a red tile within a 0-indexed grid
func parseTiles(input string) ([]grid.Coordinate, int, int) {
	rows := strings.Split(input, "\n")
	tiles, maxX, maxY := make([]grid.Coordinate, len(rows)), 0, 0
	for i, r := range rows {
		c := strings.Split(r, ",")
		x, _ := strconv.Atoi(c[0])
		y, _ := strconv.Atoi(c[1])
		tiles[i] = grid.Coordinate{X: x, Y: y}
		if x > maxX {
			maxX = x + 1
		}
		if y > maxY {
			maxY = y + 1
		}
	}
	return tiles, maxX, maxY
}

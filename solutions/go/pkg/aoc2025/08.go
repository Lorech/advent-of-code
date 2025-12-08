package aoc2025

import (
	"lorech/advent-of-code/pkg/cslices"
	"lorech/advent-of-code/pkg/grid"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Day 8: Playground
// https://adventofcode.com/2025/day/8
func dayEight(input string) (int, int) {
	return d8p1(input), 0
}

// Completes the second half of the puzzle for day 8.
func d8p1(input string, comboLimit ...int) int {
	limit := 1000
	if len(comboLimit) > 0 {
		limit = comboLimit[0]
	}

	boxes, distances := parseCoordinates(input), map[float64][]grid.Vector{}

	// Calculate the distances between each box
	for i := range len(boxes) {
		for j := range len(boxes) {
			if i == j {
				continue
			}
			a, b := boxes[i], boxes[j]
			d := straightLineDistance(a, b)
			l, exists := distances[d]
			if exists {
				if !slices.Contains(l, a) {
					distances[d] = append(l, a)
				}
				if !slices.Contains(l, b) {
					distances[d] = append(l, b)
				}
			} else {
				distances[d] = []grid.Vector{a, b}
			}
		}
	}

	// Sort all obtained distances in ascending order
	ds := make([]float64, len(distances))
	i := 0
	for k := range distances {
		ds[i] = k
		i++
	}
	slices.Sort(ds)

	// Group together all closest distance objects
	circuits, linked := make([][]grid.Vector, 0), make([]bool, len(boxes))
	for _, d := range ds[:limit] {
		a, b := distances[d][0], distances[d][1]
		ai, bi := slices.Index(boxes, a), slices.Index(boxes, b)
		al, bl := linked[ai], linked[bi]

		// If both boxes are linked, their respective circuits must be merged
		if al && bl {
			var ac, bc int
			for i, c := range circuits {
				if slices.Contains(c, a) {
					ac = i
				}
				if slices.Contains(c, b) {
					bc = i
				}
			}

			// ...except if both are already in the same circuit
			if ac == bc {
				continue
			}

			minc, maxc := ac, bc
			if bc < ac {
				minc, maxc = maxc, minc
			}
			circuits[minc] = slices.Concat(circuits[minc], circuits[maxc])
			circuits = cslices.Remove(circuits, maxc)
			continue
		}

		// If both boxes are new, this is the start of a new circuit
		if !al && !bl {
			circuits = append(circuits, []grid.Vector{a, b})
			linked[ai] = true
			linked[bi] = true
			continue
		}

		// Otherwise, link the new box to the circuit of the old box
		old, new, oldi, newi := a, b, ai, bi
		if !al && bl {
			old, new, oldi, newi = new, old, newi, oldi
		}
		for i, c := range circuits {
			if slices.Contains(c, old) {
				circuits[i] = append(circuits[i], new)
				linked[newi] = true
				break
			}
		}
	}

	slices.SortFunc(circuits, func(a, b []grid.Vector) int {
		return len(b) - len(a)
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

// Calculates the straight line or Euclidian distance between two 3D points.
// See: https://en.wikipedia.org/wiki/Euclidean_distance
// TODO: Candidate for separation into dedicated package for working in 3D.
func straightLineDistance(a, b grid.Vector) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2))
}

// Parses input data into structured data.
func parseCoordinates(input string) []grid.Vector {
	rows := strings.Split(input, "\n")
	vectors := make([]grid.Vector, len(rows))
	for i, r := range rows {
		c := strings.Split(r, ",")
		x, _ := strconv.Atoi(c[0])
		y, _ := strconv.Atoi(c[1])
		z, _ := strconv.Atoi(c[2])
		vectors[i] = grid.Vector{X: x, Y: y, Z: z}
	}
	return vectors
}

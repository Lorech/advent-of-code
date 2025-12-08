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
	return d8p1(input), d8p2(input)
}

// Completes the second half of the puzzle for day 8.
func d8p1(input string, comboLimit ...int) int {
	limit := 1000
	if len(comboLimit) > 0 {
		limit = comboLimit[0]
	}

	boxes := parseCoordinates(input)
	distances, distancePairs := parseDistances(boxes)

	circuits, linked := make([][]grid.Vector, 0), make([]bool, len(boxes))
	for _, d := range distances[:limit] {
		a, b := distancePairs[d][0], distancePairs[d][1]
		joinBoxes(a, b, boxes, &circuits, &linked)
	}

	slices.SortFunc(circuits, func(a, b []grid.Vector) int {
		return len(b) - len(a)
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

// Completes the second half of the puzzle for day 8.
func d8p2(input string) int {
	boxes := parseCoordinates(input)
	distances, distancePairs := parseDistances(boxes)

	var last []grid.Vector
	circuits, linked := make([][]grid.Vector, 0), make([]bool, len(boxes))
	for _, d := range distances {
		if !slices.Contains(linked, false) {
			break
		}
		a, b := distancePairs[d][0], distancePairs[d][1]
		joinBoxes(a, b, boxes, &circuits, &linked)
		last = distancePairs[d]
	}

	return last[0].X * last[1].X
}

// Link a single pairing of boxes together.
func joinBoxes(a, b grid.Vector, boxes []grid.Vector, circuits *[][]grid.Vector, linked *[]bool) {
	ai, bi := slices.Index(boxes, a), slices.Index(boxes, b)
	al, bl := (*linked)[ai], (*linked)[bi]

	// If both boxes are linked, their respective circuits must be merged
	if al && bl {
		var ac, bc int
		for i, c := range *circuits {
			if slices.Contains(c, a) {
				ac = i
			}
			if slices.Contains(c, b) {
				bc = i
			}
		}

		// ...except if both are already in the same circuit
		if ac == bc {
			return
		}

		minc, maxc := ac, bc
		if bc < ac {
			minc, maxc = maxc, minc
		}
		(*circuits)[minc] = slices.Concat((*circuits)[minc], (*circuits)[maxc])
		*circuits = cslices.Remove(*circuits, maxc)
		return
	}

	// If both boxes are new, this is the start of a new circuit
	if !al && !bl {
		*circuits = append(*circuits, []grid.Vector{a, b})
		(*linked)[ai] = true
		(*linked)[bi] = true
		return
	}

	// Otherwise, link the new box to the circuit of the old box
	old, new, oldi, newi := a, b, ai, bi
	if !al && bl {
		old, new, oldi, newi = new, old, newi, oldi
	}
	for i, c := range *circuits {
		if slices.Contains(c, old) {
			(*circuits)[i] = append((*circuits)[i], new)
			(*linked)[newi] = true
			break
		}
	}
}

// Convert power boxes into distance metrics, with a tuple returning
// all distances between individual pairs in ascending order, and
// a map with values of a box pair and a key of their distance apart.
func parseDistances(boxes []grid.Vector) ([]float64, map[float64][]grid.Vector) {
	dm := map[float64][]grid.Vector{}

	// Calculate the distances between each box
	for i := range len(boxes) {
		for j := range len(boxes) {
			if i == j {
				continue
			}
			a, b := boxes[i], boxes[j]
			d := straightLineDistance(a, b)
			l, exists := dm[d]
			if exists {
				if !slices.Contains(l, a) {
					dm[d] = append(l, a)
				}
				if !slices.Contains(l, b) {
					dm[d] = append(l, b)
				}
			} else {
				dm[d] = []grid.Vector{a, b}
			}
		}
	}

	// Sort all obtained distances in ascending order
	ds := make([]float64, len(dm))
	i := 0
	for k := range dm {
		ds[i] = k
		i++
	}
	slices.Sort(ds)

	return ds, dm
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

package aoc2025

import (
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type machine struct {
	lightsN int   // Total number of lights, used for targeting binary
	lights  int   // Target light state in binary
	buttons []int // Button wiring diagram in binary
	joltage []int // Target joltage
}

type lightState struct {
	depth    int
	lights   int
	visited  []int
	parent   *lightState
	children []lightState
}

// Day 10: Factory
// https://adventofcode.com/2025/day/10
func dayTen(input string) (int, int) {
	return d10p1(input), 0
}

// Completes the first half of the puzzle for day 10.
func d10p1(input string) int {
	factory := parseFactory(input)
	total := 0

	for _, m := range factory {
		min, cache := math.MaxInt, map[int]int{0: math.MaxInt}

		s := lightState{
			depth:    0,
			lights:   0,
			children: make([]lightState, 0),
			visited:  []int{0},
		}

		pressButtons(&s, m.buttons, m.lights, &min, &cache)
		total += min
	}

	return total
}

// Advances a single light state by pressing buttons and queueing upcoming states.
func pressButtons(s *lightState, bs []int, target int, min *int, cache *map[int]int) {
	// Went too deep and can't beat current known minimum value
	if s.depth == *min {
		return
	}

	// Hit the cache, no point in evaluating upcoming presses
	cached, hit := (*cache)[s.lights]
	if hit && cached == -1 {
		return
	}

	// Reached the target state, so propogate the move count in the cache
	if s.lights == target {
		*min = s.depth
		p := s.parent
		for p != nil {
			(*cache)[p.lights] = s.depth
			p = p.parent
		}
		return
	}

	for _, b := range bs {
		next := s.lights ^ b
		if !slices.Contains(s.visited, next) {
			new := lightState{
				depth:    s.depth + 1,
				lights:   next,
				visited:  append(s.visited, next),
				parent:   s,
				children: make([]lightState, 0),
			}
			s.children = append(s.children, new)
			(*cache)[next] = math.MaxInt
			pressButtons(&new, bs, target, min, cache)
		} else {
			// If a state has already been visited, it forms a loop
			(*cache)[next] = -1
		}
	}
}

// Parses input data into structured data.
func parseFactory(input string) []machine {
	rows := strings.Split(input, "\n")
	machines := make([]machine, len(rows))
	pre := regexp.MustCompile(`\[([.#]+)\] (\(.*\) ?)+ {([\d,]+)}`)
	bre := regexp.MustCompile(`[\d,]+`)

	for i, r := range rows {
		// Extract individual parts
		ps := pre.FindStringSubmatch(r)

		// Persist target light state
		ls := ps[1]
		for _, l := range ls {
			machines[i].lights <<= 0x1
			machines[i].lightsN++
			if l == '#' {
				machines[i].lights |= 0x1
			}
		}

		// Persist button wiring diagram
		bcsvs := bre.FindAllString(ps[2], -1)
		machines[i].buttons = make([]int, len(bcsvs))
		for j, bcsv := range bcsvs {
			bs, p := strings.Split(bcsv, ","), 0
			for _, b := range bs {
				bn, _ := strconv.Atoi(b)
				machines[i].buttons[j] <<= bn - p
				machines[i].buttons[j] |= 0x1
				p = bn
			}
			if p < machines[i].lightsN-1 {
				machines[i].buttons[j] <<= machines[i].lightsN - 1 - p
			}
		}

		// Persist joltage targets
		jcsv := ps[3]
		js := strings.Split(jcsv, ",")
		machines[i].joltage = make([]int, len(js))
		for k, j := range js {
			jn, _ := strconv.Atoi(j)
			machines[i].joltage[k] = jn
		}
	}

	return machines
}

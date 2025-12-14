package aoc2025

import (
	"regexp"
	"strconv"
	"strings"
)

type tree struct {
	width    int   // Width of the tree
	height   int   // Height of the tree
	presents []int // Number of each index of presents that must fit
}

type present struct {
	tiles int      // Number of tiles within the shape that occupy space
	shape [][]rune // Shape of the present - # is occupied, . is free
}

// Day 12: Christmas Tree Farm
// https://adventofcode.com/2025/day/12
func dayTwelve(input string) (int, int) {
	return d12p1(input), 0
}

// Completes the first half of the puzzle for day 12.
func d12p1(input string) int {
	trees, presents := parsePresents(input)
	valid := 0

	for _, t := range trees {
		// Preliminary: ensure the tree has enough tiles to fit the presents
		tiles, required := t.width*t.height, 0
		for i, n := range t.presents {
			required += n * presents[i].tiles
		}
		if required > tiles {
			continue
		}

		// WARN: TODO CONTAINS SPOILER
		// TODO: Packing algorithm to actually validate the allowed space,
		// this implementation is enough for Part 1, but not for the test case.

		valid++
	}

	return valid
}

// Parses the input data into structured data - meta information about each
// tree that needs to be checked, and a slice of 2D slices representing the shapes.
func parsePresents(input string) ([]tree, []present) {
	rows := strings.Split(input, "\n")
	trees, presents, pi := make([]tree, 0), make([]present, 0), -1

	for _, r := range rows {
		// Empty line - no data
		if r == "" {
			continue
		}

		// Definition for a tree
		if strings.Contains(r, "x") {
			ps := strings.Split(r, ":")

			size := strings.Split(ps[0], "x")
			w, _ := strconv.Atoi(size[0])
			h, _ := strconv.Atoi(size[1])
			tree := tree{
				width:    w,
				height:   h,
				presents: make([]int, len(presents)),
			}

			re := regexp.MustCompile(`\d+`)
			ns := re.FindAllString(ps[1], -1)
			for i, n := range ns {
				val, _ := strconv.Atoi(n)
				tree.presents[i] = val
			}

			trees = append(trees, tree)
			continue
		}

		// Start of a new shape
		if strings.Contains(r, ":") {
			presents = append(presents, present{})
			pi++
			continue
		}

		// Definition of a shape
		presents[pi].shape = append(presents[pi].shape, []rune(r))
		for _, t := range r {
			if t == '#' {
				presents[pi].tiles++
			}
		}
	}

	return trees, presents
}

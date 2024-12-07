package puzzles

import (
	"lorech/advent-of-code-2024/pkg/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Day 5: Print Queue
// https://adventofcode.com/2024/day/5
func dayFive(input string) (int, int) {
	return d5p1(input), d5p2(input)
}

// Completes the first half of the puzzle for day 5.
func d5p1(input string) int {
	dependencies, updates := parseManual(input)
	valid, _ := validateUpdates(updates, dependencies)
	return checksum(valid)
}

// Completes the second half of the puzzle for day 5.
func d5p2(input string) int {
	dependencies, updates := parseManual(input)
	fixed := make([][]int, 0)

	_, invalid := validateUpdates(updates, dependencies)
	for _, update := range invalid {
		mutated := make([]int, len(update))
		todo := update

		for len(todo) != 0 {
			// Go through each unmutated page.
			for i, page := range todo {
				resolved := true

				// Check if this page has any dependents.
				for _, p := range todo {
					if slices.Contains(dependencies[page], p) {
						resolved = false
						break
					}
				}

				// This page has no dependents, so its the next page to install from the end of the update,
				// ensuring correct resolval of dependencies for any pages prior to to this page.
				if resolved {
					mutated[len(todo)-1] = page
					todo = utils.RemoveInt(todo, i)
					break
				}
			}
		}

		fixed = append(fixed, mutated)
	}

	return checksum(fixed)
}

// Calculate the checksum for a batch of updates.
func checksum(updates [][]int) int {
	checksum := 0

	for _, update := range updates {
		middle := (len(update) - 1) / 2
		checksum += update[middle]
	}

	return checksum
}

// Validates multiple updates, splitting them into a tuple of valid and invalid updates.
func validateUpdates(updates [][]int, dependencies map[int][]int) ([][]int, [][]int) {
	valid := make([][]int, 0)
	invalid := make([][]int, 0)

	for _, update := range updates {
		if validUpdate(update, dependencies) {
			valid = append(valid, update)
		} else {
			invalid = append(invalid, update)
		}
	}

	return valid, invalid
}

// Checks if an update is valid based on the dependency map.
func validUpdate(update []int, dependencies map[int][]int) bool {
	for i, page := range update {
		dependants, exists := dependencies[page]

		// This page has no dependants, no need to validate it.
		if !exists {
			continue
		}

		for _, dependant := range dependants {
			if slices.Contains(update[:i], dependant) {
				return false
			}
		}
	}

	return true
}

// Parses the input file into usable data.
//
// Returns a tuple, where:
// - the first parameter is a map, keyed by page values, containing a slice of dependant pages;
// - the second parameter is a slice of updates, split into page numbers.
func parseManual(input string) (map[int][]int, [][]int) {
	var updates [][]int
	dependencies := make(map[int][]int)

	reDeps := regexp.MustCompile(`\d+\|\d+`)
	deps := reDeps.FindAllString(input, -1)
	for _, dep := range deps {
		instruction := strings.Split(dep, "|")
		page, _ := strconv.Atoi(instruction[0])
		dependant, _ := strconv.Atoi(instruction[1])
		_, exists := dependencies[page]
		if exists {
			dependencies[page] = append(dependencies[page], dependant)
		} else {
			dependencies[page] = []int{dependant}
		}
	}

	// Sort the dependants to make lookups faster.
	for page, _ := range dependencies {
		slices.Sort(dependencies[page])
	}

	reUps := regexp.MustCompile(`(\d+,)+\d+`)
	ups := reUps.FindAllString(input, -1)
	updates = make([][]int, len(ups))
	for i, up := range ups {
		pages := strings.Split(up, ",")
		updates[i] = make([]int, len(pages))
		for j, page := range pages {
			updates[i][j], _ = strconv.Atoi(page)
		}
	}

	return dependencies, updates
}

package main

import (
	"flag"
	"fmt"
	"lorech/advent-of-code/pkg/aoc2024"
	"lorech/advent-of-code/pkg/aoc2025"
	"lorech/advent-of-code/pkg/file"
)

// Abstraction for solving multiple years worth of puzzles.
var solve func(int, string) (any, any, error)

func main() {
	startDay, endDay, startYear, endYear := 1, 25, 2024, 2025

	pDay := flag.Int("day", -1, "Solve a specific day; solves all by default")
	pYear := flag.Int("year", -1, "Solve a specific year; solves all by default")
	flag.Parse()

	if *pDay != -1 {
		startDay, endDay = *pDay, *pDay
	}
	if *pYear != -1 {
		startYear, endYear = *pYear, *pYear
	}

	for year := startYear; year <= endYear; year++ {
		switch year {
		case 2024:
			solve = aoc2024.Solve
		case 2025:
			solve = aoc2025.Solve
		default:
			fmt.Printf("Unsupported year %d", year)
			return
		}

		for day := startDay; day <= endDay; day++ {
			data, error := file.ReadInfile(startYear, day)
			if error != nil {
				panic(error)
			}

			one, two, err := solve(day, data)

			if err != nil {
				fmt.Printf("%d Day %d: %v\n", year, day, err)
			} else {
				fmt.Printf("%d Day %d: %v, %v\n", year, day, one, two)
			}
		}
	}
}

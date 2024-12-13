package main

import (
	"flag"
	"fmt"
	"lorech/advent-of-code-2024/pkg/file"
	"lorech/advent-of-code-2024/pkg/puzzles"
)

func main() {
	start, end := 0, 13

	pDay := flag.Int("day", -1, "Solve a specific day; solves all days by default")
	flag.Parse()

	if *pDay != -1 {
		start, end = *pDay, *pDay
	}

	for day := start; day <= end; day++ {
		data, error := file.ReadInfile(day)
		if error != nil {
			panic(error)
		}

		one, two, err := puzzles.Solve(day, data)
		if err != nil {
			fmt.Printf("Day %d: %v\n", day, err)
		} else {
			fmt.Printf("Day %d: %d, %d\n", day, one, two)
		}
	}
}

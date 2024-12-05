package puzzles

// Solves the puzzle for the given day using the provided input.
//
// Layers an abstraction over the main function to allow a simpler way of
// getting the solutions of each given day, without having to implement file
// parsing within the `puzzles` package.
func Solve(day int, input string) (int, int) {
	switch day {
	case 1:
		return dayOne(input)
	case 2:
		return dayTwo(input)
	case 3:
		return dayThree(input)
	case 4:
		return dayFour(input)
	case 5:
		return dayFive(input)
	}
	return 0, 0
}

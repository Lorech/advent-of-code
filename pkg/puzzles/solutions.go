package puzzles

import "fmt"

// Solves the puzzle for the given day using the provided input.
//
// Layers an abstraction over the main function to allow a simpler way of
// getting the solutions of each given day, without having to implement file
// parsing within the `puzzles` package.
func Solve(day int, input string) (int, int, error) {
	var (
		p1, p2 int
	)
	var err error

	switch day {
	case 1:
		p1, p2 = dayOne(input)
	case 2:
		p1, p2 = dayTwo(input)
	case 3:
		p1, p2 = dayThree(input)
	case 4:
		p1, p2 = dayFour(input)
	case 5:
		p1, p2 = dayFive(input)
	case 6:
		p1, p2 = daySix(input)
	case 7:
		p1, p2 = daySeven(input)
	case 8:
		p1, p2 = dayEight(input)
	case 9:
		p1, p2 = dayNine(input)
	case 10:
		p1, p2 = dayTen(input)
	case 11:
		p1, p2 = dayEleven(input)
	case 12:
		p1, p2 = dayTwelve(input)
	case 13:
		p1, p2 = dayThirteen(input)
	case 14:
		p1, p2 = dayFourteen(input)
	case 15:
		p1, p2 = dayFifteen(input)
	default:
		err = fmt.Errorf("No implemented solution")
	}

	return p1, p2, err
}

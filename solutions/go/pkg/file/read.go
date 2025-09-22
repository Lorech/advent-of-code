package file

import (
	"fmt"
	"os"
	"strings"
)

// Reads the puzzle contents for a specific day.
//
// Expects files to be located in `infiles/{year}/{day}-{variation}.txt`
// relative to the root of the repository, where `variation` is provided in
// the third argument to the function.
func ReadInfile(year int, day int, config ...string) (string, error) {
	for i := range config {
		switch i {
		case 0:
			if config[i] != "" {
				return readFile(fmt.Sprintf("../../infiles/%04v/%02v-%v.txt", year, day, config[i]))
			}
		}
	}

	return readFile(fmt.Sprintf("../../infiles/%04v/%02v.txt", year, day))
}

// Reads the example puzzle contents for a specific day.
//
// Expects files to be located in `infiles/{year}/{day}_test-{variation}.txt`
// relative to the root of the repository, where `variation` is provided in
// the third argument to the function.
func ReadTestFile(year int, day int, config ...string) (string, error) {
	for i := range config {
		switch i {
		case 0:
			if config[i] != "" {
				return readFile(fmt.Sprintf("../../infiles/%04v/%02v_test-%v.txt", year, day, config[i]))
			}
		}
	}

	return readFile(fmt.Sprintf("../../infiles/%04v/%02v_test.txt", year, day))
}

// Reads the contents of a file relative to the root of the root of the repository.
//
// Strips the final character from the file to prevent an additional empty
// line at the end of the resulting string read into memory.
func readFile(filename string) (string, error) {
	cwd, _ := os.Getwd()
	pkg := strings.Index(cwd, "pkg")
	if pkg != -1 {
		root := cwd[:pkg]
		filename = fmt.Sprintf("%v/%v", root, filename)
	}

	data, error := os.ReadFile(filename)

	if error != nil {
		return "", error
	}

	return string(data[:len(data)-1]), nil
}

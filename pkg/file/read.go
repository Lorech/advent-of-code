package file

import (
	"fmt"
	"os"
	"strings"
)

// Reads the contents of a file, cleaning up the final newline character, and
// returns the contents of the file as a string.
//
// Helps prevent having to check for an empty line at the end of the file,
// which keeps causing issues for me.
func ReadFile(filename string) (string, error) {
	data, error := os.ReadFile(filename)

	if error != nil {
		return "", error
	}

	return string(data[:len(data)-1]), nil
}

// Reads the contents of a file for a specific day from the default input file
// directory.
//
// Expects files to be named "{day}.txt" or "{day}-{variation}.txt", where the
// variation is provided as the second parameter.
//
// By default, files will be looked for in the "infiles" directory at the root
// of the module, but setting the third parameter to `true` will look for files
// relative to the running binary, e.g., package, when running tests.
func ReadInfile(day int, config ...string) (string, error) {
	path := fmt.Sprintf("infiles/%v.txt", day)

	for i := range config {
		switch i {
		case 0:
			if config[i] != "" {
				// Add the variation if it is provided.
				path = fmt.Sprintf("infiles/%v-%v.txt", day, config[i])
			}
		case 1:
			if config[i] == "true" {
				// Force a relative path from the caller if requested.
				return ReadFile(path)
			}
		}
	}

	return readFromRoot(path)
}

// Reads the contents of a file for a specific day's example from the default
// input file directory.
//
// Expects files to be named "{day}_test.txt" or "{day}_test-{variation}.txt",
// where the variation is provided as the second parameter.
//
// By default, files will be looked for in the "infiles" directory at the root
// of the module, but setting the third parameter to `true` will look for files
// relative to the running binary, e.g., package, when running tests.
func ReadTestFile(day int, config ...string) (string, error) {
	path := fmt.Sprintf("infiles/%v_test.txt", day)

	for i := range config {
		switch i {
		case 0:
			if config[i] != "" {
				// Add the variation if it is provided.
				path = fmt.Sprintf("infiles/%v_test-%v.txt", day, config[i])
			}
		case 1:
			if config[i] == "true" {
				// Force a relative path from the caller if requested.
				return ReadFile(path)
			}
		}
	}

	return readFromRoot(path)
}

// Reads the contents of a file relative to the root of the module.
//
// When called from a main command, this should do nothing, as the file should
// already be looking inside the module root, but this helps with testing,
// allowing co-locating test data next to input data.
func readFromRoot(path string) (string, error) {
	cwd, _ := os.Getwd()
	pkg := strings.Index(cwd, "pkg")
	if pkg != -1 {
		root := cwd[:pkg]
		path = fmt.Sprintf("%v/%v", root, path)
	}

	return ReadFile(path)
}

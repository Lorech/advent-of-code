package benchmarks

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Process() {
	files := findUnprocessedFiles()
	bms := make([]Benchmark, 0)

	for _, file := range files {
		d, error := os.ReadFile(file)
		if error != nil {
			fmt.Printf("Warning: Could not read file %v. Reason: %v", file, error)
			continue
		}

		bm, err := processFile(string(d))
		if err != nil {
			fmt.Printf("Warning: Could not process file %v. Reason: %v", file, error)
			continue
		}

		bms = append(bms, bm)
	}

	// TODO: Compare the new data to the existing data.

	// TODO: Update the benchmark history file.

	// TODO: Update the new benchmarks to avoid processing them again.
}

// Finds all unprocessed benchmark files within the benchmark directory.
func findUnprocessedFiles() []string {
	files := make([]string, 0)

	content, err := os.ReadDir("benchmarks")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range content {
		if !f.IsDir() {
			if !strings.Contains(f.Name(), "-processed") {
				files = append(files, fmt.Sprintf("benchmarks/%s", f.Name()))
			}
		}
	}

	return files
}

// Processes a benchmark file into a usable struct for further processing.
func processFile(data string) (Benchmark, error) {
	bm := Benchmark{}

	// Validate that the benchmark is in regular TXT output.
	// TODO: There is probably a better way to check this.
	valid := strings.Contains(data, "goos: ")
	if !valid {
		return bm, fmt.Errorf("File is not a valid benchmark TXT file.")
	}

	rows := strings.Split(data, "\n")
	bm.results = make([]BenchmarkResult, len(rows)-6)

	for i, row := range rows {
		if strings.HasPrefix(row, "Benchmark") {
			r := strings.Fields(row)
			f := r[0]
			c, _ := strconv.Atoi(r[1])
			l, _ := strconv.Atoi(r[2])
			bm.results[i-4] = BenchmarkResult{f, c, l}
			continue
		}

		if strings.HasPrefix(row, "goos:") {
			s := strings.Index(row, " ")
			bm.os = row[s+1:]
			continue
		}

		if strings.HasPrefix(row, "goarch:") {
			s := strings.Index(row, " ")
			bm.arch = row[s+1:]
			continue
		}

		if strings.HasPrefix(row, "pkg:") {
			s := strings.Index(row, " ")
			bm.pkg = row[s+1:]
			continue
		}

		if strings.HasPrefix(row, "cpu:") {
			s := strings.Index(row, " ")
			bm.cpu = row[s+1:]
			continue
		}
	}

	return bm, nil
}

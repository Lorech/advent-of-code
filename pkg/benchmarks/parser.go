package benchmarks

import (
	"fmt"
	"log"
	"lorech/advent-of-code-2024/pkg/cmaps"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Parses any new benchmark files and compares them with the previous bests,
// adding any new benchmarks or updating new bests as necessary.
func Process() {
	files := findUnprocessedFiles()
	pFiles := make([]string, 0)
	bms := make([]Benchmark, 0)

	for _, file := range files {
		d, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Warning: Could not read file %v. Reason: %v", file, err)
			continue
		}

		bm, err := processFile(string(d))
		if err != nil {
			fmt.Printf("Warning: Could not process file %v. Reason: %v", file, err)
			continue
		}

		bms = append(bms, bm)
		pFiles = append(pFiles, file)
	}

	pb, _ := PersonalBests()
	pb = updatePersonalBest(pb, bms)
	savePersonalBests(pb)
	updateProcessedFiles(pFiles)
}

// Merges existing personal bests with a batch of new benchmark data into
// once complete set of data.
// WARNING: This currently assumes an existing and identical set of metadata
// for a benchmark, only updating functions. This is subject to change.
func updatePersonalBest(pb Benchmark, benchmarks []Benchmark) Benchmark {
	for _, benchmark := range benchmarks {
		pb = mergeBenchmarks(pb, benchmark)
	}
	return pb
}

// Combine two benchmarks, returning a single benchmark with the best results
// for each function across the two of them.
// WARNING: This currently assumes an existing and identical set of metadata
// for a benchmark, only updating functions. This is subject to change.
func mergeBenchmarks(a Benchmark, b Benchmark) Benchmark {
	aFuncs := cmaps.KeysSlice(a.Results)
	bFuncs := cmaps.KeysSlice(b.Results)

	// Update the runtimes of the functions in a.
	for f, result := range a.Results {
		if slices.Contains(bFuncs, f) {
			if result.Ops < b.Results[f].Ops && result.Ns < b.Results[f].Ns {
				a.Results[f] = b.Results[f]
			}
		}
	}

	// Set all the meta data from b to a, as b is expected to be more up-to-date.
	if a.Pkg == "" {
		a.Pkg = b.Pkg
	}
	if a.Os == "" {
		a.Os = b.Os
	}
	if a.Arch == "" {
		a.Arch = b.Arch
	}
	if a.Cpu == "" {
		a.Cpu = b.Cpu
	}

	// Add any missing functions to a, as b is expected to be more up-to-date.
	for f, result := range b.Results {
		if !slices.Contains(aFuncs, f) {
			if a.Results == nil {
				a.Results = make(map[string]BenchmarkResult)
			}
			a.Results[f] = result
		}
	}

	return a
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
	bm.Results = make(map[string]BenchmarkResult, len(rows)-6)

	for _, row := range rows {
		if strings.HasPrefix(row, "Benchmark") {
			r := strings.Fields(row)
			f := r[0]
			c, _ := strconv.Atoi(r[1])
			l, _ := strconv.Atoi(r[2])
			bm.Results[f] = BenchmarkResult{f, c, l}
			continue
		}

		if strings.HasPrefix(row, "goos:") {
			s := strings.Index(row, " ")
			bm.Os = row[s+1:]
			continue
		}

		if strings.HasPrefix(row, "goarch:") {
			s := strings.Index(row, " ")
			bm.Arch = row[s+1:]
			continue
		}

		if strings.HasPrefix(row, "pkg:") {
			s := strings.Index(row, " ")
			bm.Pkg = row[s+1:]
			continue
		}

		if strings.HasPrefix(row, "cpu:") {
			s := strings.Index(row, " ")
			bm.Cpu = row[s+1:]
			continue
		}
	}

	return bm, nil
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
			if strings.Contains(f.Name(), "bests.json") {
				continue
			}

			if !strings.Contains(f.Name(), "-processed") {
				files = append(files, fmt.Sprintf("benchmarks/%s", f.Name()))
			}
		}
	}

	return files
}

// Renames all processed files to avoid processing them again on future runs.
func updateProcessedFiles(files []string) {
	for _, file := range files {
		ext := strings.Index(file, ".")
		if ext != 1 {
			os.Rename(file, fmt.Sprintf("%s-processed%s", file[:ext], file[ext:]))
		} else {
			os.Rename(file, fmt.Sprintf("%s-processed", file))
		}
	}
}

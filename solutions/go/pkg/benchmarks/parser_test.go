package benchmarks

import (
	"os"
	"testing"
)

// Checks file lookup from the benchmarks directory, ensuring processed files
// get skipped.
func TestFileLookup(t *testing.T) {
	e := []string{"benchmarks/sample.txt"}
	r := findUnprocessedFiles()

	if len(r) != len(e) || r[0] != e[0] {
		t.Errorf("findUnprocessedFiles() = %v, expected %v", r, e)
	}
}

// Checks if a benchmark output file can be properly parsed into a struct.
func TestFileProcessing(t *testing.T) {
	d, error := os.ReadFile("benchmarks/sample.txt")
	if error != nil {
		t.Fatalf("Could not read file: Reason: %v", error)
	}

	e := Benchmark{
		"lorech/advent-of-code-2024/pkg/benchmarks",
		"darwin",
		"arm64",
		"Apple M1 Pro",
		[]BenchmarkResult{
			{
				"BenchmarkNumberOne-8",
				3656,
				328111,
			},
			{
				"BenchmarkNumberTwo-8",
				3758,
				3225689,
			},
			{
				"BenchmarkNumberThree-8",
				1148,
				1031164,
			},
		},
	}

	r, err := processFile(string(d))
	if err != nil {
		t.Fatal(err)
	}

	// TODO: Test this more thoroughly.
	if r.pkg != e.pkg {
		t.Fatalf("processFile() = %v, expected %v", r, e)
	}
}

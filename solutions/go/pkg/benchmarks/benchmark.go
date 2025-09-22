package benchmarks

import (
	"encoding/json"
	"fmt"
	"os"
)

// The output of a single benchmark file.
type Benchmark struct {
	Pkg     string                     `json:"pkg"`     // The package the benchmarks were for.
	Os      string                     `json:"os"`      // The OS used to run the benchmarks.
	Arch    string                     `json:"arch"`    // The architecture used to run the benchmarks.
	Cpu     string                     `json:"cpu"`     // The CPU used to run the benchmarks.
	Results map[string]BenchmarkResult `json:"results"` // The results of this benchmark.
}

// The output of a single benchmark.
type BenchmarkResult struct {
	Fun string `json:"fun"` // The name of the function tested. This should match the key of the map used to access this.
	Ops int    `json:"ops"` // The cycles taken to benchmark the function.
	Ns  int    `json:"ns"`  // The time of a single benchmark cycle, in nanoseconds.
}

func (a *BenchmarkResult) Equal(b *BenchmarkResult) bool {
	return a.Fun == b.Fun && a.Ns == b.Ns
}

// Fetches the personal best benchmarks from the stored file on disk.
func PersonalBests() (Benchmark, error) {
	var b Benchmark

	d, err := os.ReadFile("benchmarks/bests.json")
	if err != nil {
		fmt.Printf("Warning: Could not read records file. Reason: %v\n", err)
		return b, err
	}

	err = json.Unmarshal(d, &b)
	if err != nil {
		fmt.Printf("Warning: Could not process records file. Reason: %v\n", err)
	}

	return b, err
}

// Stores a benchmark into the personal bests file.
func savePersonalBests(data Benchmark) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Errorf("Error: Could not encode records file. Reason :%v\n", err)
		return
	}

	err = os.WriteFile("benchmarks/bests.json", b, 0666)
	if err != nil {
		fmt.Errorf("Error: Could not write records file. Reason: %v\n", err)
	}
}

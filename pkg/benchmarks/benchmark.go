package benchmarks

// The output of a single benchmark file.
type Benchmark struct {
	pkg     string            // The package the benchmarks were for.
	os      string            // The OS used to run the benchmarks.
	arch    string            // The architecture used to run the benchmarks.
	cpu     string            // The CPU used to run the benchmarks.
	results []BenchmarkResult // The results of this benchmark.
}

// The output of a single benchmark.
type BenchmarkResult struct {
	fun string // The name of the function tested.
	ops int    // The cycles taken to benchmark the function.
	ns  int    // The time of a single benchmark cycle, in nanoseconds.
}

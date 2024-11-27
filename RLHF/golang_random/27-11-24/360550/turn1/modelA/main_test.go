package main

import (
	"testing"
)

// Example function to be benchmarked
func sumNumbers() int {
	var total int
	for i := 0; i < 1_000_000; i++ {
		total += i
	}
	return total
}

// Benchmark function
func BenchmarkSumNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumNumbers()
	}
}

func main() {
	// Run the benchmark
	testing.Benchmark(BenchmarkSumNumbers)
}

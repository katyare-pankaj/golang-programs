package main

import (
	"testing"
)

// The function you want to benchmark
func slowFunction(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

// Benchmark function using the "testing" package
func BenchmarkSlowFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slowFunction(1000000)
	}
}

func main() {
	// Run the benchmark
	testing.Benchmark(BenchmarkSlowFunction)
}

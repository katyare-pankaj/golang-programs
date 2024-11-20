package main

import (
	"testing"
)

// Factorial calculates the factorial of a non-negative integer n.
func Factorial(n int) int {
	if n < 0 {
		panic("Factorial is not defined for negative numbers")
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// BenchmarkFactorial benchmarks the Factorial function.
func BenchmarkFactorial(b *testing.B) {
	n := 10
	for i := 0; i < b.N; i++ {
		Factorial(n)
	}
}

func main() {
	// Run the benchmark
	testing.Benchmark(BenchmarkFactorial)

	// You can also run profiling to identify bottlenecks
	// go tool pprof main -prof cpu
	// go tool pprof -web cpu
}

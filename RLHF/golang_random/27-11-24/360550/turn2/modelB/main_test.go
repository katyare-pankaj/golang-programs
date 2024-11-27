package main

import (
	"testing"
)

func Factorial(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func BenchmarkFactorial(b *testing.B) {
	// Test input for benchmarking.
	var input uint64 = 20

	// Run the benchmark loop for the specified number of times (b.N).
	for i := 0; i < b.N; i++ {
		Factorial(input)
	}
}

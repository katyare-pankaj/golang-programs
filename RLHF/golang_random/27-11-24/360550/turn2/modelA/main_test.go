package main

import (
	"testing"
)

// Function to be benchmarked
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Benchmark function
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(20) // Benchmarking fibonacci with n = 20
	}
}

func TestFibonacci(t *testing.T) {
	if fibonacci(10) != 55 {
		t.Error("Fibonacci function returned incorrect result")
	}
}

func main() {
	testing.Main()
}

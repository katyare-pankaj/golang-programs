package main

import (
	"testing"
	"time"
)

// Original function with a simulated delay
func inefficientFunction(data []int) int {
	sum := 0
	for _, value := range data {
		// Simulated heavy calculation
		time.Sleep(10 * time.Millisecond)
		sum += value
	}
	return sum
}

// Optimized function without the delay
func efficientFunction(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}

// Benchmark test for inefficientFunction with memory allocation reporting
func BenchmarkInefficientFunction(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ReportAllocs() // Enable memory allocation reporting
	for i := 0; i < b.N; i++ {
		inefficientFunction(data)
	}
}

// Benchmark test for efficientFunction with memory allocation reporting
func BenchmarkEfficientFunction(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ReportAllocs() // Enable memory allocation reporting
	for i := 0; i < b.N; i++ {
		efficientFunction(data)
	}
}

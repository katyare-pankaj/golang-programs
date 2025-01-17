package main

import (
	"fmt"
	"testing"
	"time"
)

// The function to be benchmarked
func BenchmarkSomeFunction(b *testing.B) {
	// Your function logic here
}

// Function to perform the warm-up iterations
func Warmup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSomeFunction(b)
	}
}

func BenchmarkAdaptiveWarmup(b *testing.B) {
	minWarmup := 10
	maxWarmup := 100
	stepWarmup := 10

	warmupIterations := minWarmup
	avgDuration := time.Duration(0)
	numIterations := 3 // Number of iterations to average

	// Feedback loop to adjust warm-up iterations
	for i := 0; i < 5; i++ { // Number of rounds of feedback
		b.ResetTimer()
		Warmup(b)
		b.StopTimer()

		dur := time.Duration(b.Nanoseconds()) / time.Duration(warmupIterations)
		fmt.Printf("Warmup Iterations: %d, Average Duration: %s\n", warmupIterations, dur)

		if dur > time.Millisecond*50 { // Adjust the threshold as needed
			// If average duration is high, increase warm-up iterations
			warmupIterations += stepWarmup
			warmupIterations = min(warmupIterations, maxWarmup)
		} else if dur < time.Millisecond*20 {
			// If average duration is low, decrease warm-up iterations
			warmupIterations -= stepWarmup
			warmupIterations = max(warmupIterations, minWarmup)
		} else {
			// If average duration is within the range, we're done
			break
		}
	}

	// Run the benchmark with the final warm-up iterations
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BenchmarkSomeFunction(b)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

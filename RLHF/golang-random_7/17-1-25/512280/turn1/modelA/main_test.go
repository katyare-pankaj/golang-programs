package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// TestFunction represents the function we're benchmarking.
func TestFunction() {
	// Simulate some work using sleep
	time.Sleep(10 * time.Millisecond)
}

// benchmarkFunction runs the provided function `fn` up to `b.N` times.
func benchmarkFunction(b *testing.B, fn func()) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn()
	}
}

// adaptiveWarmUp determines the optimal number of warm-up iterations before steady
// performance is reached.
func adaptiveWarmUp(b *testing.B, fn func()) int {
	const (
		targetRelativeChange = 0.05 // Target change in execution time (5%)
		maxWarmUps           = 100  // Maximum number of warm-up iterations
	)

	var (
		prevTime    time.Duration
		currentTime time.Duration
		changes     []float64
	)

	for warmUpIterations := 1; warmUpIterations <= maxWarmUps; warmUpIterations++ {
		start := time.Now()
		fn()
		end := time.Now()

		currentTime = end.Sub(start)

		if warmUpIterations > 1 {
			change := math.Abs(float64(currentTime-prevTime)) / float64(prevTime)
			changes = append(changes, change)

			// If the average change over the last three iterations is below the target,
			// stop warming up.
			if len(changes) > 3 {
				avgChange := (changes[len(changes)-1] + changes[len(changes)-2] + changes[len(changes)-3]) / 3
				if avgChange < targetRelativeChange {
					return warmUpIterations
				}
			}
		}

		prevTime = currentTime
	}

	// Return maximum if steady state not reached within the limit
	return maxWarmUps
}

func BenchmarkAdaptiveFunction(b *testing.B) {
	warmUpCount := adaptiveWarmUp(b, TestFunction)
	fmt.Printf("Determined warm-up iterations: %d\n", warmUpCount)

	// Execute the warm-up iterations
	for i := 0; i < warmUpCount; i++ {
		TestFunction()
	}

	// Perform the actual benchmark
	benchmarkFunction(b, TestFunction)
}

func main() {
	// Calling the benchmark test manually (without `go test` tool)
	result := testing.Benchmark(BenchmarkAdaptiveFunction)
	fmt.Printf("%s\n", result)
}

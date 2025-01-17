package main

import (
	"fmt"
	"math"
	"runtime"
	"testing"
	"time"
)

// TestFunction represents the function we're benchmarking.
func TestFunction() {
	// Simulate some work using sleep
	time.Sleep(time.Millisecond * 10)
}

// benchmarkFunction runs the provided function fn up to b.N times.
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
		minWarmUps           = 5    // Minimum number of warm-up iterations
	)
	var (
		prevTime         time.Duration
		currentTime      time.Duration
		changes          []float64
		warmUpIterations = minWarmUps
	)
	for ; warmUpIterations <= maxWarmUps; warmUpIterations++ {
		start := time.Now()
		fn()
		end := time.Now()
		currentTime = end.Sub(start)
		if warmUpIterations > minWarmUps {
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
		// Adjust the warm-up iteration count based on the execution time.
		// If the current iteration took longer, reduce the number of iterations to
		// focus on faster iterations in the next warm-up phase.
		if currentTime > 5*time.Millisecond {
			warmUpIterations--
		}
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
	runtime.GC()
	// Perform the actual benchmark
	benchmarkFunction(b, TestFunction)
}

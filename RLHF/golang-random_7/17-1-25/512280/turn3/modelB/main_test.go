package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// ExampleFunction is the function we want to benchmark.
func ExampleFunction() {
	// Simulate some work. Replace this with your actual code.
	time.Sleep(10 * time.Millisecond)
}

func adaptiveWarmUp(b *testing.B, fn func()) {
	const (
		targetStability = 0.05 // Target stability percentage (5%)
		maxWarmUps      = 1000 // Maximum number of warm-up iterations
		minWarmUps      = 2    // Minimum number of warm-up iterations
		windowSize      = 5    // Number of iterations to observe for stability
	)

	var (
		prevTime   time.Duration
		currTime   time.Duration
		changes    []float64
		numWarmUps int
	)

	// Execute at least minWarmUps iterations to get an initial baseline
	for numWarmUps < minWarmUps {
		start := time.Now()
		fn()
		currTime = time.Since(start)
		prevTime = currTime
		numWarmUps++
	}

	for ; numWarmUps <= maxWarmUps; numWarmUps++ {
		start := time.Now()
		fn()
		currTime = time.Since(start)

		change := math.Abs(float64(currTime-prevTime)) / float64(prevTime)
		changes = append(changes, change)

		if len(changes) == windowSize {
			maxChange := math.Max(changes[0], changes[1], changes[2], changes[3], changes[4])
			minChange := math.Min(changes[0], changes[1], changes[2], changes[3], changes[4])

			if (maxChange-minChange)/minChange < targetStability {
				break
			}

			changes = changes[1:]
		}

		prevTime = currTime
	}

	b.SetBytes(int64(numWarmUps))
}

func BenchmarkAdaptiveFunction(b *testing.B) {
	adaptiveWarmUp(b, ExampleFunction)

	// Perform the actual benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExampleFunction()
	}
}

func main() {
	result := testing.Benchmark(BenchmarkAdaptiveFunction)
	fmt.Printf("%s\n", result)
}

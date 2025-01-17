package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// Function to be benchmarked
func exampleFunction() {
	// Simulate work; replace with actual logic
	time.Sleep(10 * time.Millisecond)
}

// adaptiveWarmUp executes the function repeatedly and stops when the execution time
// stabilizes within a set threshold.
func adaptiveWarmUp(fn func()) int {
	const (
		targetRelativeChange = 0.01 // Target change rate in execution time (1%)
		maxWarmUps           = 100  // Maximum warm-up iterations allowed
		stabilityWindow      = 3    // Number of iterations to observe for stability
	)

	var (
		prevTime  time.Duration
		currTime  time.Duration
		changes   []float64
		numWarmUp int
	)

	for numWarmUp = 1; numWarmUp <= maxWarmUps; numWarmUp++ {
		start := time.Now()
		fn()
		currTime = time.Since(start)

		if numWarmUp > 1 {
			change := math.Abs(float64(currTime-prevTime)) / float64(prevTime)
			changes = append(changes, change)

			// Check if the average change over the stability window is below the threshold
			if len(changes) >= stabilityWindow {
				avgChange := 0.0
				for _, ch := range changes[len(changes)-stabilityWindow:] {
					avgChange += ch
				}
				avgChange /= float64(stabilityWindow)

				if avgChange < targetRelativeChange {
					return numWarmUp
				}
			}
		}

		prevTime = currTime
	}

	return maxWarmUps
}

// BenchmarkAdaptiveFunction benchmarks the function using an adaptive warm-up phase.
func BenchmarkAdaptiveFunction(b *testing.B) {
	warmUpCount := adaptiveWarmUp(exampleFunction)
	fmt.Printf("Determined warm-up iterations: %d\n", warmUpCount)

	// Execute warm-up iterations
	for i := 0; i < warmUpCount; i++ {
		exampleFunction()
	}

	// Perform the actual benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		exampleFunction()
	}
}

func main() {
	// Simulate running the benchmark through testing.Benchmark without the `go test` tool
	result := testing.Benchmark(BenchmarkAdaptiveFunction)
	fmt.Printf("%s\n", result)
}

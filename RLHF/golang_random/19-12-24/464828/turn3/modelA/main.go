package main

import (
	"fmt"
	"math"
)

// Function without named return values
func calculateStatsWithoutNamedReturn(numbers []float64) (float64, float64) {
	sum := 0.0
	var variance float64
	for _, num := range numbers {
		sum += num
	}
	mean := sum / float64(len(numbers))

	for _, num := range numbers {
		diff := num - mean
		variance += diff * diff
	}
	variance /= float64(len(numbers))
	stdDev := math.Sqrt(variance)

	return mean, stdDev
}

// Function with named return values
func calculateStatsWithNamedReturn(numbers []float64) (mean, stdDev float64) {
	sum := 0.0
	var variance float64
	for _, num := range numbers {
		sum += num
	}
	mean = sum / float64(len(numbers))

	for _, num := range numbers {
		diff := num - mean
		variance += diff * diff
	}
	variance /= float64(len(numbers))
	stdDev = math.Sqrt(variance)

	return
}

func main() {
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println("Using unnamed return values:")
	meanWithout, stdDevWithout := calculateStatsWithoutNamedReturn(numbers)
	fmt.Printf("Mean: %.2f, Std Dev: %.2f\n", meanWithout, stdDevWithout)

	fmt.Println("\nUsing named return values:")
	meanWith, stdDevWith := calculateStatsWithNamedReturn(numbers)
	fmt.Printf("Mean: %.2f, Std Dev: %.2f\n", meanWith, stdDevWith)
}

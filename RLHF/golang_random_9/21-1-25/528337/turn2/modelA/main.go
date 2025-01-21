package main

import (
	"fmt"
	"math"
	"sort"
)

// Mean calculates the mean of a slice of numbers.
func Mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Median calculates the median of a slice of numbers.
// Note: The original data slice is not modified.
func Median(data []float64) float64 {
	n := len(data)
	if n == 0 {
		panic("cannot calculate median of an empty slice")
	}

	// Create a copy of the data to sort
	sortedData := append([]float64(nil), data...)
	sort.Float64s(sortedData)

	// Calculate median
	if n%2 == 1 {
		return sortedData[n/2]
	}
	return (sortedData[n/2-1] + sortedData[n/2]) / 2
}

// StandardDeviation calculates the standard deviation of a slice of numbers.
func StandardDeviation(data []float64) float64 {
	if len(data) == 0 {
		panic("cannot calculate standard deviation of an empty slice")
	}

	mean := Mean(data)
	var sumSquares float64
	for _, value := range data {
		sumSquares += math.Pow(value-mean, 2)
	}
	variance := sumSquares / float64(len(data))
	return math.Sqrt(variance)
}

func main() {
	// Original immutable data slice
	data := []float64{10, 20, 30, 40, 50}

	// Calculate mean, median, and standard deviation
	mean := Mean(data)
	median := Median(data)
	stdDev := StandardDeviation(data)

	// Output the results
	fmt.Printf("Data: %v\n", data)
	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Median: %.2f\n", median)
	fmt.Printf("Standard Deviation: %.2f\n", stdDev)

	// Check that the original data has not been modified
	fmt.Printf("Original data (unchanged): %v\n", data)
}

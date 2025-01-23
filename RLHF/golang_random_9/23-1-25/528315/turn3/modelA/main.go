package main

import (
	"fmt"
	"math"
)

// normalize scales the values in a slice to a 0-1 range, handling edge cases.
func normalize(data []float64) []float64 {
	n := len(data)
	if n == 0 {
		// Return an empty slice if the dataset is empty
		return []float64{}
	}

	min, max := minMax(data)

	normalized := make([]float64, n)

	// Handle the uniform data case to avoid division by zero
	if max == min {
		// Return a slice filled with a constant value, e.g., all zeros or another relevant value
		for i := 0; i < n; i++ {
			normalized[i] = 0
		}
		return normalized
	}

	// Standard normalization process
	rangeVal := max - min
	for i := 0; i < n; i++ {
		normalized[i] = (data[i] - min) / rangeVal
	}

	return normalized
}

// minMax finds the minimum and maximum values in a slice.
func minMax(data []float64) (min, max float64) {
	min = math.Inf(1)
	max = math.Inf(-1)
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	// Examples to demonstrate normalization
	datasets := [][]float64{
		{3.0, 6.0, 9.0, 12.0, 15.0},
		{5.0, 5.0, 5.0}, // uniformly the same elements
		{},              // empty dataset
	}

	for i, data := range datasets {
		fmt.Printf("Dataset %d: %v\n", i+1, data)
		fmt.Printf("Normalized %d: %v\n\n", i+1, normalize(data))
	}
}

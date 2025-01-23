package main

import (
	"fmt"
	"math"
)

// normalize efficiently normalizes the dataset slice to a 0-1 range.
func normalize(data []float64) []float64 {
	// Handle edge case for empty datasets
	if len(data) == 0 {
		return []float64{}
	}

	min, max := minMax(data)

	// Create a new slice with the same size to store normalized values
	normalized := make([]float64, len(data))

	// Check if all values are the same to avoid division by zero
	if max-min == 0 {
		for i := range data {
			normalized[i] = 0.5 // Or any predefined constant, as they are all the same
		}
		return normalized
	}

	// Normalize each data point
	for i, value := range data {
		normalized[i] = (value - min) / (max - min)
	}

	return normalized
}

// minMax returns the minimum and maximum values in a dataset slice.
func minMax(data []float64) (min, max float64) {
	// Initialize min and max
	min = math.Inf(1)
	max = math.Inf(-1)

	for _, value := range data {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}

func main() {
	// Test datasets
	datasets := [][]float64{
		{3.0, 6.0, 9.0, 12.0, 15.0},
		{5.0, 5.0, 5.0}, // uniform dataset
		{},              // empty dataset
	}

	for _, data := range datasets {
		normalizedData := normalize(data)
		fmt.Println("Original Data:", data)
		fmt.Println("Normalized Data:", normalizedData)
		fmt.Println()
	}
}

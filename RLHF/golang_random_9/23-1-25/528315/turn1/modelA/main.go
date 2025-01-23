package main

import (
	"fmt"
	"math"
)

// normalize normalizes the values in the dataset to a 0-1 range.
func normalize(data []float64) []float64 {
	if len(data) == 0 {
		return nil
	}

	min, max := minMax(data)

	// Create a new slice to store normalized values
	normalized := make([]float64, len(data))

	for i, value := range data {
		if max-min != 0 {
			normalized[i] = (value - min) / (max - min)
		} else {
			normalized[i] = 0 // Avoid division by zero if all values are the same
		}
	}

	return normalized
}

// minMax finds the minimum and maximum values in a dataset.
func minMax(data []float64) (min, max float64) {
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
	// Example dataset
	data := []float64{3.0, 6.0, 9.0, 12.0, 15.0}

	// Normalize the dataset
	normalizedData := normalize(data)

	// Display the normalized dataset
	fmt.Println("Normalized Data:", normalizedData)
}

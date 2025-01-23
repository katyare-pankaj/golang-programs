package main

import (
	"fmt"
)

// Normalize normalizes the values in the dataset to a 0-1 range.
func Normalize(data []float64) []float64 {
	if len(data) == 0 {
		// Return an empty slice for an empty dataset
		return nil
	}

	if len(data) == 1 {
		// Return the same slice for a single-element dataset
		return data
	}

	min, max := MinMax(data)

	// Create a new slice to store the normalized values with the same capacity as the input slice
	normalized := make([]float64, len(data), len(data))

	for i, value := range data {
		normalized[i] = (value - min) / (max - min)
	}

	return normalized
}

// MinMax finds the minimum and maximum values in the dataset.
func MinMax(data []float64) (min, max float64) {
	min = data[0]
	max = data[0]

	for _, value := range data[1:] {
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}

	return
}

func main() {
	// Example dataset
	data := []float64{3.0, 6.0, 9.0, 12.0, 15.0}
	normalizedData := Normalize(data)
	fmt.Println("Normalized Data:", normalizedData)

	// Example of an empty dataset
	emptyDataset := []float64{}
	normalizedEmptyDataset := Normalize(emptyDataset)
	fmt.Println("Normalized Empty Dataset:", normalizedEmptyDataset)

	// Example of a single-element dataset
	singleElementDataset := []float64{1.0}
	normalizedSingleElementDataset := Normalize(singleElementDataset)
	fmt.Println("Normalized Single-Element Dataset:", normalizedSingleElementDataset)
}

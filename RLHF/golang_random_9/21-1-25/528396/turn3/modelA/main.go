package main

import (
	"fmt"
)

// Function to calculate the mode
func mode(data []float64) []float64 {
	// Create a frequency map to count occurrences of each number
	frequencyMap := make(map[float64]int)
	for _, value := range data {
		frequencyMap[value]++
	}

	// Determine the maximum frequency
	maxFrequency := 0
	for _, frequency := range frequencyMap {
		if frequency > maxFrequency {
			maxFrequency = frequency
		}
	}

	// Collect all numbers with the maximum frequency
	var modes []float64
	for number, frequency := range frequencyMap {
		if frequency == maxFrequency {
			modes = append(modes, number)
		}
	}

	return modes
}

func main() {
	// Original data slice (immutable)
	originalData := []float64{4.2, 2.7, 4.2, 8.1, 2.7, 4.2, 3.9}

	// Calculate the mode
	modes := mode(originalData)

	fmt.Println("Original data:", originalData)
	fmt.Println("Mode(s):", modes)

	// Confirm that the original data remains unchanged
	fmt.Println("Original data after calculations:", originalData)
}

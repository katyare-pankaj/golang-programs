package main

import (
	"fmt"
)

func calculateAverage(numbers []float64) float64 {
	// Check for empty slice to avoid division by zero
	if len(numbers) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func main() {
	testCases := [][]float64{
		{1.0, 2.0, 3.0},
		{}, // Empty slice test
		{42.0},
	}

	for _, numbers := range testCases {
		avg := calculateAverage(numbers)
		fmt.Println("Average of", numbers, ":", avg)
	}
}

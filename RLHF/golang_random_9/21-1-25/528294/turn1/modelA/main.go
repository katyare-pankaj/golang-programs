package main

import (
	"fmt"
	"math"
)

// calculateMean computes the mean of the slice, ignoring NaN values.
func calculateMean(data []float64) float64 {
	var sum float64
	count := 0
	for _, value := range data {
		if !math.IsNaN(value) {
			sum += value
			count++
		}
	}
	if count == 0 {
		return math.NaN() // Return NaN if there are no valid numbers
	}
	return sum / float64(count)
}

// imputeWithMean replaces NaN values with the mean.
func imputeWithMean(data []float64) []float64 {
	mean := calculateMean(data)
	for i, value := range data {
		if math.IsNaN(value) {
			data[i] = mean
		}
	}
	return data
}

// removeMissingValues removes entries with NaN values.
func removeMissingValues(data []float64) []float64 {
	cleanedData := []float64{}
	for _, value := range data {
		if !math.IsNaN(value) {
			cleanedData = append(cleanedData, value)
		}
	}
	return cleanedData
}

func main() {
	data := []float64{1.0, 2.0, math.NaN(), 4.0, 5.0, math.NaN(), 7.0}

	fmt.Println("Original data:", data)

	// Option 1: Impute with mean
	imputedData := imputeWithMean(make([]float64, len(data)))
	copy(imputedData, data)
	imputedData = imputeWithMean(imputedData)
	fmt.Println("Data after imputing with mean:", imputedData)

	// Option 2: Remove entries with missing values
	cleanedData := removeMissingValues(data)
	fmt.Println("Data after removing missing values:", cleanedData)
}

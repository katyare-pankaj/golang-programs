package main

import (
	"fmt"
)

func main() {
	// Initialize a slice with some missing values (represented by -1)
	data := []float64{2.5, 3.6, 4.8, -1, 7.2, 5.4, -1, 9.0, 6.3}

	// Function to calculate the mean of the valid data
	mean := func(values []float64) float64 {
		var sum float64
		var count int
		for _, value := range values {
			if value != -1 {
				sum += value
				count++
			}
		}
		return sum / float64(count)
	}

	// Calculate the mean of the valid data
	validMean := mean(data)
	fmt.Println("Mean of valid data:", validMean)

	// Imputation: Replace missing values with the mean
	imputedData := make([]float64, len(data))
	copy(imputedData, data)
	for i, value := range imputedData {
		if value == -1 {
			imputedData[i] = validMean
		}
	}
	fmt.Println("Data with missing values imputed:", imputedData)

	// Removal: Create a new slice without missing values
	filteredData := []float64{}
	for _, value := range data {
		if value != -1 {
			filteredData = append(filteredData, value)
		}
	}
	fmt.Println("Data with missing values removed:", filteredData)
}

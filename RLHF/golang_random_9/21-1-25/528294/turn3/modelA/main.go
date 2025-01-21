package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []float64{1.0, 2.0, -999.0, 4.0, -999.0, 6.0, 7.0, 8.0, -999.0, 10.0}
	fmt.Println("Original Data: ", data)
	median := calculateMedian(data)
	fmt.Println("Median of Valid Data: ", median)
	imputedData := replaceMissingWithMedian(data, median)
	fmt.Println("Imputed Data: ", imputedData)
}

func calculateMedian(data []float64) float64 {
	// Filter out the missing values (-999)
	var validData []float64
	for _, value := range data {
		if value != -999.0 {
			validData = append(validData, value)
		}
	}
	// Sort the valid data to find the median
	sort.Float64s(validData)

	n := len(validData)
	if n == 0 {
		return 0 // Handle case where there are no valid data points
	}

	// Calculate median
	if n%2 == 0 {
		// Even number of elements
		return (validData[n/2-1] + validData[n/2]) / 2
	}
	// Odd number of elements
	return validData[n/2]
}

func replaceMissingWithMedian(data []float64, median float64) []float64 {
	imputedData := make([]float64, len(data))
	copy(imputedData, data)
	for i, value := range imputedData {
		if value == -999.0 {
			imputedData[i] = median
		}
	}
	return imputedData
}

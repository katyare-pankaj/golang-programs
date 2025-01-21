package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []float64{1.0, 2.0, -999.0, 4.0, -999.0, 6.0, 7.0, 8.0, -999.0, 10.0, 3.0}
	fmt.Println("Original Data: ", data)
	median := calculateMedian(data)
	fmt.Println("Median of Valid Data: ", median)
	imputedData := replaceMissingWithMedian(data, median)
	fmt.Println("Imputed Data: ", imputedData)
}

func calculateMedian(data []float64) float64 {
	// Sort the valid data
	validData := []float64{}
	for _, value := range data {
		if value != -999.0 {
			validData = append(validData, value)
		}
	}
	sort.Float64s(validData)

	// Calculate the median
	length := len(validData)
	if length == 0 {
		return 0
	}

	if length%2 == 0 {
		return (validData[length/2-1] + validData[length/2]) / 2
	} else {
		return validData[length/2]
	}
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

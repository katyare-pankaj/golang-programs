package main

import (
	"fmt"
)

func main() {
	data := []float64{1.0, 2.0, -999.0, 4.0, -999.0, 6.0, 7.0, 8.0, -999.0, 10.0}
	fmt.Println("Original Data: ", data)
	mean := calculateMean(data)
	fmt.Println("Mean of Valid Data: ", mean)
	imputedData := replaceMissingWithMean(data, mean)
	fmt.Println("Imputed Data: ", imputedData)
}

func calculateMean(data []float64) float64 {
	var sum float64
	var count int
	for _, value := range data {
		if value != -999.0 {
			sum += value
			count++
		}
	}
	return sum / float64(count)
}

func replaceMissingWithMean(data []float64, mean float64) []float64 {
	imputedData := make([]float64, len(data))
	copy(imputedData, data)
	for i, value := range imputedData {
		if value == -999.0 {
			imputedData[i] = mean
		}
	}
	return imputedData
}

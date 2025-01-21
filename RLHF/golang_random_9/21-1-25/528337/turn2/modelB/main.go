package main

import (
	"fmt"
	"math"
	"sort"
)

func mean(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

func median(data []float64) float64 {
	sortedData := append([]float64(nil), data...) // Create a copy of the slice
	sort.Float64s(sortedData)
	n := len(sortedData)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		midIndex1 := n/2 - 1
		midIndex2 := n / 2
		return (sortedData[midIndex1] + sortedData[midIndex2]) / 2
	} else {
		midIndex := n / 2
		return sortedData[midIndex]
	}
}

func standardDeviation(data []float64) float64 {
	avg := mean(data)
	var sumSquaredDifferences float64
	for _, v := range data {
		sumSquaredDifferences += (v - avg) * (v - avg)
	}
	return math.Sqrt(sumSquaredDifferences / float64(len(data)))
}

func main() {
	// Original data slice (immutable)
	originalData := []float64{2.3, 4.5, 7.8, 1.2, 3.6, 5.4}

	fmt.Println("Original data:", originalData)

	// Calculate mean, median, and standard deviation
	meanValue := mean(originalData)
	medianValue := median(originalData)
	stdDevValue := standardDeviation(originalData)

	fmt.Println("Mean:", meanValue)
	fmt.Println("Median:", medianValue)
	fmt.Println("Standard Deviation:", stdDevValue)

	// Confirm that the original data remains unchanged
	fmt.Println("Original data after calculations:", originalData)
}

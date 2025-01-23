package main

import (
	"fmt"
	"math"
	"sync"
)

func normalize(data []float64) []float64 {
	// Handle edge case for empty datasets
	if len(data) == 0 {
		return []float64{}
	}

	// Find min and max values
	min, max := minMax(data)

	// If all values are the same (min == max), return a slice of zeros
	if min == max {
		return make([]float64, len(data))
	}

	// Normalize the data in parallel using goroutines
	normalized := make([]float64, len(data))
	var wg sync.WaitGroup
	wg.Add(len(data))
	for i := range data {
		go func(i int) {
			normalized[i] = (data[i] - min) / (max - min)
			wg.Done()
		}(i)
	}
	wg.Wait()

	return normalized
}

func minMax(data []float64) (min, max float64) {
	min = math.Inf(1)
	max = math.Inf(-1)

	// Use channels to compute min and max in parallel
	minCh := make(chan float64, 1)
	maxCh := make(chan float64, 1)
	go func() {
		var localMin float64 = math.Inf(1)
		for _, value := range data {
			if value < localMin {
				localMin = value
			}
		}
		minCh <- localMin
	}()
	go func() {
		var localMax float64 = math.Inf(-1)
		for _, value := range data {
			if value > localMax {
				localMax = value
			}
		}
		maxCh <- localMax
	}()

	min = <-minCh
	max = <-maxCh

	return
}

func main() {
	// Test datasets
	datasets := [][]float64{
		{3.0, 6.0, 9.0, 12.0, 15.0},
		{5.0, 5.0, 5.0}, // Uniform dataset
		{},              // Empty dataset
	}
	for _, data := range datasets {
		normalizedData := normalize(data)
		fmt.Println("Original Data:", data)
		fmt.Println("Normalized Data:", normalizedData)
		fmt.Println()
	}
}

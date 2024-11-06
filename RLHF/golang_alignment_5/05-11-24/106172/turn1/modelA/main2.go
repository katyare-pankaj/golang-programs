package main

import (
	"fmt"
	"math"
	"time"
)

// HFTAlgorithm represents the HFT trading algorithm
type HFTAlgorithm struct {
	windowSize int
	prices     []float64
	sum        float64
}

// NewHFTAlgorithm creates a new HFTAlgorithm instance
func NewHFTAlgorithm(windowSize int) *HFTAlgorithm {
	return &HFTAlgorithm{
		windowSize: windowSize,
		prices:     make([]float64, 0, windowSize),
	}
}

// CalculateMovingAverage calculates the moving average of the current prices
func (a *HFTAlgorithm) CalculateMovingAverage() float64 {
	if len(a.prices) < a.windowSize {
		return 0
	}

	a.sum -= a.prices[0]
	copy(a.prices, a.prices[1:])
	return a.sum / float64(a.windowSize)
}

// UpdatePrice updates the algorithm with the latest price
func (a *HFTAlgorithm) UpdatePrice(price float64) {
	a.sum += price
	a.prices = append(a.prices, price)
	if len(a.prices) > a.windowSize {
		a.sum -= a.prices[0]
		a.prices = a.prices[1:]
	}
}

// ... (Previous code remains the same)

func evaluateMAPerformance(windowSize int, prices []float64) float64 {
	algorithm := NewHFTAlgorithm(windowSize)
	var totalError float64

	for _, price := range prices {
		algorithm.UpdatePrice(price)
		ma := algorithm.CalculateMovingAverage()
		// Simple error calculation for demonstration
		// Assuming we want to minimize the absolute difference from the simple average
		simpleAvg := sum(prices) / float64(len(prices))
		error := math.Abs(ma - simpleAvg)
		totalError += error
	}
	return totalError
}

func sum(numbers []float64) float64 {
	total := 0.0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	prices := []float64{100, 105, 102, 108, 103, 110, 107, 104, 106, 109, 112}
	// Initial window size
	currentWindowSize := 10
	bestWindowSize := currentWindowSize
	minError := evaluateMAPerformance(currentWindowSize, prices)

	fmt.Printf("Initial Window Size: %d, Error: %.2f\n", currentWindowSize, minError)

	// Scrum Iterative Tuning Loop
	for sprint := 1; sprint <= 3; sprint++ {
		fmt.Printf("\n-- Sprint %d --\n", sprint)
		for _, windowSize := range []int{8, 12, 15, 20} { // Possible window sizes to try
			start := time.Now()
			error := evaluateMAPerformance(windowSize, prices)
			duration := time.Since(start)
			fmt.Printf("Window Size: %d, Error: %.2f, Time: %s\n", windowSize, error, duration)

			if error < minError {
				minError = error
				bestWindowSize = windowSize
				fmt.Println("New best window size found!")
			}
		}
		// Refine the range based on the results of the previous sprint
		// For simplicity, let's just adjust the start and end by 2 in this example
		startWindowSize := bestWindowSize - 2
		endWindowSize := bestWindowSize + 2

		fmt.Printf("Best Window Size so far: %d, Range for next sprint: [%d, %d]\n", bestWindowSize, startWindowSize, endWindowSize)
	}

	fmt.Println("\n-- Optimization Complete --")
	fmt.Printf("Best Window Size: %d, Min Error: %.2f\n", bestWindowSize, minError)
}

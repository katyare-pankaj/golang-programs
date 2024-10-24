package main

import (
	"fmt"
)

// Interval represents a closed interval [min, max]
type Interval struct {
	Min float64
	Max float64
}

// analyzeMovingAverage calculates the interval of possible values for the moving average given the intervals of input data points and window size
func analyzeMovingAverage(dataInterval Interval, windowSize int) Interval {
	if windowSize <= 0 {
		panic("Window size must be positive")
	}

	// Assuming data points are non-negative
	minAverage := dataInterval.Min / float64(windowSize)
	maxAverage := dataInterval.Max / float64(windowSize)

	return Interval{Min: minAverage, Max: maxAverage}
}

func main() {
	// Example usage:
	dataInterval := Interval{Min: 0.0, Max: 100.0} // Assuming data points are in the range [0, 100]
	windowSize := 5
	resultInterval := analyzeMovingAverage(dataInterval, windowSize)

	fmt.Printf("Interval of possible moving averages: [%.2f, %.2f]\n", resultInterval.Min, resultInterval.Max)

	// Real-time data processing simulation
	dataStream := []float64{23.4, 45.6, 67.8, 89.0, 54.3}
	movingAverage := calculateMovingAverage(dataStream, windowSize)
	fmt.Printf("Current Moving Average: %.2f\n", movingAverage)
}

func calculateMovingAverage(data []float64, windowSize int) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

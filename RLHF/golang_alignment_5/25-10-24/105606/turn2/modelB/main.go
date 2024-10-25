package main

import (
	"fmt"
	"math"
	"sort"
)

func isAnomaly(value float64, dataset []float64) bool {
	// Calculate z-score
	sortedData := append([]float64{}, dataset...)
	sort.Float64s(sortedData)
	mean := calculateMean(sortedData)
	stdDev := calculateStdDev(sortedData)
	zScore := (value - mean) / stdDev
	// Set an anomaly threshold (e.g., z-score > 3)
	threshold := 3.0
	return math.Abs(zScore) > threshold
}
func calculateMean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}
func calculateStdDev(data []float64) float64 {
	mean := calculateMean(data)
	squaredDifferences := make([]float64, len(data))
	for i, value := range data {
		squaredDifferences[i] = math.Pow(value-mean, 2)
	}
	variance := calculateMean(squaredDifferences)
	return math.Sqrt(variance)
}

func main() {
	// Log data points (e.g., response times)
	logData := []float64{100.0, 120.0, 80.0, 150.0, 110.0, 130.0, 140.0, 180.0, 160.0, 200.0}

	anomalyValue := 300.0 // Potential anomaly
	isAnomaly := isAnomaly(anomalyValue, logData)
	fmt.Println("Is Anomaly? :", isAnomaly)
}

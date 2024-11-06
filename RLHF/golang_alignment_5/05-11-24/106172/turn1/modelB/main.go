package main

import (
	"fmt"
	"math"
)

func calculateSMA(prices []float64, windowSize int) float64 {
	sum := 0.0
	for _, price := range prices {
		sum += price
	}
	return sum / float64(windowSize)
}

func optimizeSMAWindowSize(prices []float64, minWindowSize, maxWindowSize int) (int, float64) {
	bestWindowSize := minWindowSize
	bestSharpeRatio := -math.MaxFloat64

	for windowSize := minWindowSize; windowSize <= maxWindowSize; windowSize++ {
		smaShort := calculateSMA(prices[0:windowSize], windowSize)
		smaLong := calculateSMA(prices[windowSize:2*windowSize], windowSize)
		// Calculate Sharpe ratio for the current window size
		sharpeRatio := calculateSharpeRatio(prices, smaShort, smaLong)
		if sharpeRatio > bestSharpeRatio {
			bestWindowSize = windowSize
			bestSharpeRatio = sharpeRatio
		}
	}
	return bestWindowSize, bestSharpeRatio
}

func main() {
	// Sample market data for the stock
	prices := []float64{100, 105, 102, 108, 103, 110, 107, 104, 106, 112}

	// Define the range of SMA window sizes to optimize
	minWindowSize := 20
	maxWindowSize := 50

	// Optimize the SMA window size
	bestWindowSize, bestSharpeRatio := optimizeSMAWindowSize(prices, minWindowSize, maxWindowSize)

	fmt.Printf("Best SMA window size: %d\n", bestWindowSize)
	fmt.Printf("Best Sharpe ratio: %f\n", bestSharpeRatio)
}

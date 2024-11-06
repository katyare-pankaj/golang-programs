package main

import (
	"fmt"
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

func main() {
	// Initial window size assumption
	windowSize := 10
	algorithm := NewHFTAlgorithm(windowSize)

	// Simulate price updates
	prices := []float64{100, 105, 102, 108, 103, 110, 107, 104, 106, 109, 112}

	for _, price := range prices {
		algorithm.UpdatePrice(price)
		ma := algorithm.CalculateMovingAverage()
		fmt.Printf("Price: %.2f, MA: %.2f\n", price, ma)
	}
}

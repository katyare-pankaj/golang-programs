package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gonum/stat"
)

const (
	windowSize = 10 // Size of the sliding window for monitoring
)

func main() {
	// Generate random data representing CPU usage percentages over time
	rand.Seed(time.Now().UnixNano())
	cpuUsageStream := make(chan float64)

	go func() {
		for {
			time.Sleep(time.Second)
			cpuUsageStream <- rand.Float64() * 100
		}
	}()

	monitorAndOptimize(cpuUsageStream)
}

func monitorAndOptimize(cpuUsageStream <-chan float64) {
	var slidingWindow []float64
	var modelPredictions []float64

	for cpuUsage := range cpuUsageStream {
		// Add the current CPU usage to the sliding window
		slidingWindow = append(slidingWindow, cpuUsage)

		// Keep the window size fixed
		if len(slidingWindow) > windowSize {
			slidingWindow = slidingWindow[1:]
		}

		// Wait for enough data to populate the window
		if len(slidingWindow) < windowSize {
			continue
		}

		// Fit a simple linear regression model to the sliding window data
		m, b := fitLinearRegression(slidingWindow, modelPredictions)

		// Predict the CPU usage for the next time step
		nextPrediction := m*float64(len(slidingWindow)) + b
		modelPredictions = append(modelPredictions, nextPrediction)

		// Analyze the prediction and take action if needed
		if nextPrediction > 90 { // Threshold for high CPU usage
			fmt.Println("High CPU usage predicted! Taking action...")
			// Add your optimization logic here, such as scaling resources,
			// adjusting task priorities, or triggering auto-scaling mechanisms.
			// For simplicity, we'll just print a message.
		}

		// Display current status
		fmt.Printf("Current CPU Usage: %.2f%%\tPredicted Next: %.2f%%\n", cpuUsage, nextPrediction)
	}
}

// fitLinearRegression fits a linear regression model to the given data.
func fitLinearRegression(x, y []float64) (m, b float64) {
	// Calculate mean of x and y
	meanX := stat.Mean(x, nil)
	meanY := stat.Mean(y, nil)

	// Calculate covariance and variance of x and y
	cov := stat.Covariance(x, y, nil)
	varX := stat.Variance(x, nil)

	// Calculate slope (m) and intercept (b)
	m = cov / varX
	b = meanY - m*meanX

	return
}

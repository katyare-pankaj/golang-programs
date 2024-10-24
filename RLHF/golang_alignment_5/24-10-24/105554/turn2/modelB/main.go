package main

import (
	"fmt"
	"time"
)

// TemperatureEvent represents a temperature data event
type TemperatureEvent struct {
	Timestamp   time.Time
	Temperature float64
}

// Interval represents a closed interval [min, max]
type Interval struct {
	Min float64
	Max float64
}

// analyzeTemperatureInterval analyzes the interval of possible temperature values based on historical data
func analyzeTemperatureInterval(historicalData []TemperatureEvent) Interval {
	// Assuming historical data is available
	minTemperature := historicalData[0].Temperature
	maxTemperature := historicalData[0].Temperature

	for _, event := range historicalData[1:] {
		if event.Temperature < minTemperature {
			minTemperature = event.Temperature
		} else if event.Temperature > maxTemperature {
			maxTemperature = event.Temperature
		}
	}

	// Add a small buffer to the interval for fluctuations
	buffer := 1.0
	return Interval{Min: minTemperature - buffer, Max: maxTemperature + buffer}
}

// processTemperatureEvent processes a temperature event using interval analysis
func processTemperatureEvent(event TemperatureEvent, expectedInterval Interval) {
	if event.Temperature < expectedInterval.Min || event.Temperature > expectedInterval.Max {
		// Event value falls outside the expected interval, take appropriate action
		// For example, trigger an alert or adjust the processing pipeline
		fmt.Printf("Temperature event out of interval: %.2f\n", event.Temperature)
	} else {
		// Event value falls within the expected interval, apply optimized processing
		// For example, perform further calculations or send the event to the next stage of processing
		fmt.Printf("Processing temperature event: %.2f\n", event.Temperature)
	}
}

func main() {
	// Simulate historical data collection
	historicalData := []TemperatureEvent{
		{Timestamp: time.Now(), Temperature: 25.0},
		{Timestamp: time.Now().Add(time.Minute), Temperature: 28.0},
		{Timestamp: time.Now().Add(2 * time.Minute), Temperature: 22.0},
		// Add more historical data points
	}

	// Analyze the interval of possible temperature values
	expectedInterval := analyzeTemperatureInterval(historicalData)
	fmt.Printf("Expected temperature interval: [%.2f, %.2f]\n", expectedInterval.Min, expectedInterval.Max)

	// Simulate streaming data
	streamingData := []TemperatureEvent{
		{Timestamp: time.Now().Add(3 * time.Minute), Temperature: 24.0},
		{Timestamp: time.Now().Add(4 * time.Minute), Temperature: 30.0},
		{Timestamp: time.Now().Add(5 * time.Minute), Temperature: 18.0},
		// Add more streaming data points
	}

	// Process the streaming data events using interval analysis
	for _, event := range streamingData {
		processTemperatureEvent(event, expectedInterval)
	}
}

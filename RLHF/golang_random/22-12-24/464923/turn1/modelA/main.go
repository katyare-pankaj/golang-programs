package main

import (
	"log"
	"sync"
	"time"
)

// Simulate a data point
type DataPoint struct {
	Value float64
}

// Preprocess function for scaling
func scaleData(data []DataPoint) []DataPoint {
	log.Println("Scaling data...")
	time.Sleep(time.Duration(500) * time.Millisecond)
	for i := range data {
		data[i].Value *= 2.0
	}
	return data
}

// Preprocess function for standardization
func standardizeData(data []DataPoint) []DataPoint {
	log.Println("Standardizing data...")
	time.Sleep(time.Duration(700) * time.Millisecond)
	for i := range data {
		data[i].Value -= 10.0
	}
	return data
}

// Preprocess function for encoding (in this case, a simple no-op placeholder)
func encodeData(data []DataPoint) []DataPoint {
	log.Println("Encoding data...")
	time.Sleep(time.Duration(300) * time.Millisecond)
	return data
}

// Main function to simulate the machine learning pipeline
func main() {
	// Sample data for preprocessing
	data := []DataPoint{{1.0}, {2.0}, {3.0}, {4.0}, {5.0}}

	// Create a WaitGroup to keep track of concurrent goroutines
	var wg sync.WaitGroup

	// Launch the preprocessing tasks in goroutines
	log.Println("Starting preprocessing steps...")

	wg.Add(1) // Start scaling in a goroutine
	go func() {
		defer wg.Done()
		data = scaleData(data)
	}()

	wg.Add(1) // Start standardizing in a goroutine
	go func() {
		defer wg.Done()
		data = standardizeData(data)
	}()

	wg.Add(1) // Start encoding in a goroutine
	go func() {
		defer wg.Done()
		data = encodeData(data)
	}()

	// Wait for all the goroutines to finish
	wg.Wait()

	log.Println("Preprocessing steps completed.")
	log.Printf("Final data: %+v\n", data)
}

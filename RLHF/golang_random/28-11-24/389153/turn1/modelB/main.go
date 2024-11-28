package main

import (
	"fmt"
	"log"
)

// Define the interface for data processing
type DataProcessor interface {
	ProcessData() error
}

// Define a struct to represent historical data source (e.g., database)
type HistoricalDataProcessor struct {
	// Add any necessary fields for historical data processing
}

// Implement the ProcessData method for historical data processing
func (h *HistoricalDataProcessor) ProcessData() error {
	// Fetch data from the database and process it
	fmt.Println("Processing historical data...")
	// Placeholder for database fetch and processing logic
	return nil
}

// Define a struct to represent real-time data source (e.g., Kafka stream)
type RealTimeDataProcessor struct {
	// Add any necessary fields for real-time data processing
}

// Implement the ProcessData method for real-time data processing
func (r *RealTimeDataProcessor) ProcessData() error {
	// Consume data from the Kafka stream and process it
	fmt.Println("Processing real-time data...")
	// Placeholder for Kafka consumption and processing logic
	return nil
}

// This function uses the DataProcessor interface to process data from any source
func ProcessData(processor DataProcessor) error {
	return processor.ProcessData()
}

func main() {
	// Create instances of historical and real-time data processors
	historicalProcessor := &HistoricalDataProcessor{}
	realTimeProcessor := &RealTimeDataProcessor{}

	// Use the DataProcessor interface to process data from different sources
	if err := ProcessData(historicalProcessor); err != nil {
		log.Fatal(err)
	}

	if err := ProcessData(realTimeProcessor); err != nil {
		log.Fatal(err)
	}
}

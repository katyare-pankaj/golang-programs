package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Simulate different stages of a data processing pipeline
func ingestData(data chan<- []string, wg *sync.WaitGroup) {
	log.Println("Starting data ingestion...")
	time.Sleep(time.Duration(1000) * time.Millisecond)
	data <- []string{"Data 1", "Data 2", "Data 3", "Data 4", "Data 5"}
	log.Println("Data ingestion complete.")
	defer wg.Done()
}

func cleanData(data <-chan []string, cleanedData chan<- []string, wg *sync.WaitGroup) {
	log.Println("Starting data cleaning...")
	time.Sleep(time.Duration(1500) * time.Millisecond)
	cleanedData <- []string{"Cleaned Data 1", "Cleaned Data 2", "Cleaned Data 3", "Cleaned Data 4", "Cleaned Data 5"}
	log.Println("Data cleaning complete.")
	defer wg.Done()
}

func transformData(cleanedData <-chan []string, transformedData chan<- []string, wg *sync.WaitGroup) {
	log.Println("Starting data transformation...")
	time.Sleep(time.Duration(2000) * time.Millisecond)
	transformedData <- []string{"Transformed Data 1", "Transformed Data 2", "Transformed Data 3", "Transformed Data 4", "Transformed Data 5"}
	log.Println("Data transformation complete.")
	defer wg.Done()
}

func exportData(transformedData <-chan []string, wg *sync.WaitGroup) {
	log.Println("Starting data export...")
	time.Sleep(time.Duration(2500) * time.Millisecond)
	receivedData := <-transformedData
	for _, data := range receivedData {
		fmt.Println(data)
	}
	log.Println("Data export complete.")
	defer wg.Done()
}

func main() {
	// Create channels to pass data between stages
	data := make(chan []string)
	cleanedData := make(chan []string)
	transformedData := make(chan []string)

	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Start goroutines for each task in the pipeline
	wg.Add(1)
	go ingestData(data, &wg)

	wg.Add(1)
	go cleanData(data, cleanedData, &wg)

	wg.Add(1)
	go transformData(cleanedData, transformedData, &wg)

	wg.Add(1)
	go exportData(transformedData, &wg)

	// Wait for all goroutines to complete
	wg.Wait()

	log.Println("Pipeline execution complete.")
}

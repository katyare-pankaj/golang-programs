package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	// Process the CSV file
	processCSVFile("large_data.csv")
}

func processCSVFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Create a reader for the CSV file
	reader := csv.NewReader(file)

	// Store data in memory using a sync.Map
	dataMap := &sync.Map{}

	// Read the CSV file and store data in the sync.Map
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading record: %v", err)
		}

		// Assuming the CSV file has two columns: key and value
		key := record[0]
		value := record[1]

		dataMap.Store(key, value)
	}

	// Perform analysis on the data (using the dataMap)
	// ...

	// Clear the dataMap to free up memory
	dataMap.Range(func(key, _ interface{}) bool {
		dataMap.Delete(key)
		return true
	})
}

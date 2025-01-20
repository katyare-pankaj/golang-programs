package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	// Open the CSV file
	file, err := os.Open("example.csv")
	if err != nil {
		log.Fatalf("Failed to open the CSV file: %s\n", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Read all fields per record

	// Allocate memory for records
	var records [][]string
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println("Initial memory allocation:", mem.Alloc)
	// Read the CSV data into records
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		records = append(records, record)
	}

	// Print the CSV data
	for i, record := range records {
		fmt.Println(i, record)
	}

	// Check memory usage after reading CSV data
	runtime.ReadMemStats(&mem)
	fmt.Println("Memory allocation after reading CSV data:", mem.Alloc)
}

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read data from CSV file
	inputFile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Perform data transformation (uppercasing each string)
	transformedRecords := make([][]string, len(records))
	for i, record := range records {
		transformedRecord := make([]string, len(record))
		for j, val := range record {
			transformedRecord[j] = strings.ToUpper(val)
		}
		transformedRecords[i] = transformedRecord
	}

	// Write transformed data to another CSV file
	outputFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	err = writer.WriteAll(transformedRecords)
	if err != nil {
		log.Fatalf("Error writing to output file: %v", err)
	}

	fmt.Println("Data processing completed successfully!")
}

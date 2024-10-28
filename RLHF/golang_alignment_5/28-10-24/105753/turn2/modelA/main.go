package main

import (
	"fmt"
	"strings"
)

// Step 1: Define the DataProcessor struct
type DataProcessor struct{}

// Step 2: Refactor the processData function into smaller, descriptive methods
func (dp *DataProcessor) ProcessData(inputData []string) []string {
	data := dp.extractData(inputData)
	data = dp.transformData(data)
	data = dp.loadData(data)
	return data
}

func (dp *DataProcessor) extractData(inputData []string) []string {
	// Simulate data extraction
	return inputData
}

func (dp *DataProcessor) transformData(extractedData []string) []string {
	cleanedData := make([]string, 0)
	for _, record := range extractedData {
		cleanedRecord := dp.cleanRecord(record)
		if cleanedRecord != "" {
			cleanedData = append(cleanedData, cleanedRecord)
		}
	}
	return cleanedData
}

func (dp *DataProcessor) cleanRecord(record string) string {
	// Perform data cleaning operations
	return strings.TrimSpace(strings.ToLower(record))
}

func (dp *DataProcessor) loadData(transformedData []string) []string {
	// Simulate data loading
	return transformedData
}

func main() {
	inputData := []string{
		" Apple ",
		"Banana",
		"Cherry ",
		"",
		"dApple",
	}

	processor := &DataProcessor{}
	processedData := processor.ProcessData(inputData)

	for _, processedRecord := range processedData {
		fmt.Println(processedRecord)
	}
}

package main

import (
	"fmt"
	"strings"
)

type DataCleaner struct {
}

func (dc *DataCleaner) CleanData(inputData []string) []string {
	return dc.transformData(dc.extractData(inputData))
}

func (dc *DataCleaner) extractData(inputData []string) []string {
	// Simulate data extraction
	return inputData
}

func (dc *DataCleaner) transformData(extractedData []string) []string {
	cleanedData := make([]string, 0)
	for _, data := range extractedData {
		cleanedData = append(cleanedData, dc.cleanRecord(data))
	}
	return cleanedData
}

func (dc *DataCleaner) cleanRecord(record string) string {
	// Perform data cleaning operations here
	return strings.TrimSpace(strings.ToLower(record))
}

func main() {
	inputData := []string{
		" Apple ",
		"Banana",
		"Cherry ",
		"",
		"dApple",
	}

	cleaner := &DataCleaner{}
	cleanedData := cleaner.CleanData(inputData)

	for _, cleanedRecord := range cleanedData {
		fmt.Println(cleanedRecord)
	}
}

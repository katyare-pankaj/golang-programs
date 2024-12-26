package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func processDataset(dataset []string) {
	for index, entry := range dataset {
		if entry == "" {
			// Handle empty entry
			log.Printf("Skipping empty entry at index %d", index)
			continue
		}

		// Simulate processing of the entry
		result, err := processEntry(entry)
		if err != nil {
			// Log error
			log.Printf("Error processing entry at index %d: %v", index, err)
			continue
		}

		// Log successful processing
		log.Printf("Successfully processed entry at index %d: %s", index, result)
	}
}

func processEntry(entry string) (string, error) {
	// Simulate some processing that might fail
	parts := strings.Split(entry, ",")
	if len(parts) != 2 {
		return "", fmt.Errorf("entry '%s' is not in the correct format", entry)
	}

	name, ageStr := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return "", fmt.Errorf("cannot convert age '%s' to integer: %v", ageStr, err)
	}

	// Successfully processed the entry
	return fmt.Sprintf("Processed: %s, %d", name, age), nil
}

func main() {
	dataset := []string{
		"Alice, 30",
		"Bob, 25",
		"Charlie", // Missing age
		"David, ", // Missing name
		"Eve, 35",
	}

	processDataset(dataset)
}

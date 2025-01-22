package main

import (
	"fmt"
	"strings"
	"testing"
)

// Sample data to simulate extraction
var sourceData = []string{
	"john doe",
	"jane smith",
	"emily jones",
}

// transformData applies a transformation to each element.
// In this case, we capitalize each name.
func transformData(data []string) []string {
	var transformedData []string
	for _, name := range data {
		// Transform: Capitalize first letter of each word
		transformedName := strings.Title(name)
		transformedData = append(transformedData, transformedName)
	}
	return transformedData
}

// loadData simulates loading data to a destination
func loadData(data []string) {
	fmt.Println("Loading data...")
	for _, record := range data {
		fmt.Println(record)
	}
}

func main() {
	// Extract
	fmt.Println("Extracting data...")
	extractedData := sourceData

	// Transform
	fmt.Println("Transforming data...")
	transformedData := transformData(extractedData)

	// Load
	loadData(transformedData)
}

// Test function to validate transformation logic
func TestTransformData(t *testing.T) {
	input := []string{"john doe", "jane smith", "emily jones"}
	expectedOutput := []string{"John Doe", "Jane Smith", "Emily Jones"}

	result := transformData(input)

	if len(result) != len(expectedOutput) {
		t.Errorf("Expected result length %d, but got %d", len(expectedOutput), len(result))
	}

	for i, v := range result {
		if v != expectedOutput[i] {
			t.Errorf("Expected %s, but got %s", expectedOutput[i], v)
		}
	}
}

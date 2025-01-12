package main

import (
	"fmt"
)

// generateFormattedStrings generates a sequence of strings formatted with a pattern and given bounds.
func generateFormattedStrings(pattern string, start int, end int) []string {
	if start > end {
		return []string{}
	}

	strings := make([]string, 0, end-start+1) // Initialize the slice with a good capacity to avoid reallocations.

	for i := start; i <= end; i++ {
		str := fmt.Sprintf(pattern, i)
		strings = append(strings, str)
	}

	return strings
}

func main() {
	// Pattern defines the string format: "Result%03d"
	pattern := "Result%03d"
	start := 1
	end := 10

	// Generate the sequence of formatted strings
	result := generateFormattedStrings(pattern, start, end)

	// Print each formatted string for clarity
	fmt.Println("Formatted Sequence:")
	for _, str := range result {
		fmt.Println(str)
	}
}

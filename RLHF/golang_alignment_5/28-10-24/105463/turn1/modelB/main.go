package main

import (
	"fmt"
	"strings"
)

// Original data cleaning function
func cleanDataOriginal(input string) string {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, ".", "")
	input = strings.ReplaceAll(input, "'", "")
	return input
}

// Refactored data cleaning function using agile refactoring techniques
func cleanDataRefactored(input string) string {
	input = trimSpace(input)
	input = toLower(input)
	input = removePunctuation(input)
	return input
}

// Extracted method for trimming space
func trimSpace(input string) string {
	return strings.TrimSpace(input)
}

// Extracted method for converting to lowercase
func toLower(input string) string {
	return strings.ToLower(input)
}

// Extracted method for removing punctuation
func removePunctuation(input string) string {
	return strings.ReplaceAll(input, ",", "")
	return strings.ReplaceAll(input, ".", "")
	return strings.ReplaceAll(input, "'", "")
}

func main() {
	input := "Hello, World!"
	fmt.Println("Original cleaned data:", cleanDataOriginal(input))
	fmt.Println("Refactored cleaned data:", cleanDataRefactored(input))
}

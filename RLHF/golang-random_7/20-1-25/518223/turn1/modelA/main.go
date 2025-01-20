package main

import (
	"fmt"
	"strings"
)

// Define a function type for callback
type StringProcessor func(string) string

// ProcessStrings processes a slice of strings using a callback
func ProcessStrings(strings []string, processor StringProcessor) []string {
	var result []string
	for _, str := range strings {
		processedStr := processor(str)
		result = append(result, processedStr)
	}
	return result
}

// UpperCaseProcessor is a callback function that converts a string to uppercase
func UpperCaseProcessor(s string) string {
	return strings.ToUpper(s)
}

// ReverseProcessor is a callback function that reverses a string
func ReverseProcessor(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	strings := []string{"hello", "world", "go", "language"}

	// Use UpperCaseProcessor as a callback
	uppercasedStrings := ProcessStrings(strings, UpperCaseProcessor)
	fmt.Println("Uppercase Strings:", uppercasedStrings)

	// Use ReverseProcessor as a callback
	reversedStrings := ProcessStrings(strings, ReverseProcessor)
	fmt.Println("Reversed Strings:", reversedStrings)

	// Use an anonymous function as a callback
	customProcessedStrings := ProcessStrings(strings, func(s string) string {
		return "[" + s + "]"
	})
	fmt.Println("Custom Processed Strings:", customProcessedStrings)
}

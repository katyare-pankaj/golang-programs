package main

import (
	"fmt"
	"strings"
)

// StringProcessor is a function type that takes a string and returns a string.
type StringProcessor func(string) string

// ProcessStrings applies a sequence of processors to each string in the slice.
func ProcessStrings(strings []string, processors ...StringProcessor) []string {
	var result []string
	for _, str := range strings {
		var processedStr = str
		for _, processor := range processors {
			processedStr = processor(processedStr)
		}
		result = append(result, processedStr)
	}
	return result
}

// UpperCaseProcessor converts a string to uppercase.
func UpperCaseProcessor(s string) string {
	return strings.ToUpper(s)
}

// ReverseProcessor reverses a string.
func ReverseProcessor(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CustomFormatter formats a string with a specified prefix and suffix.
func CustomFormatter(prefix, suffix string) StringProcessor {
	return func(s string) string {
		return prefix + s + suffix
	}
}

func main() {
	strings := []string{"hello", "world", "go", "language"}

	// Simple usage: Applying a single processor
	uppercasedStrings := ProcessStrings(strings, UpperCaseProcessor)
	fmt.Println("Uppercase Strings:", uppercasedStrings)

	// Creating a custom processor and applying it
	boldFormatter := CustomFormatter("<b>", "</b>")
	boldStrings := ProcessStrings(strings, boldFormatter)
	fmt.Println("Bold Strings:", boldStrings)

	// Creating a processing pipeline with multiple processors
	pipeline := []StringProcessor{
		UpperCaseProcessor,
		ReverseProcessor,
		CustomFormatter("(", ")"),
	}
	processedStrings := ProcessStrings(strings, pipeline...)
	fmt.Println("Processed Strings (pipeline):", processedStrings)

	// Using an anonymous function as a callback
	customProcessedStrings := ProcessStrings(strings, func(s string) string {
		return strings.ReplaceAll(s, "o", "0")
	})
	fmt.Println("Custom Processed Strings:", customProcessedStrings)
}

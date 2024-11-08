package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Map function: Takes a line of text and returns a list of word-count pairs
func mapper(line string) []map[string]int {
	words := strings.Fields(line)
	result := []map[string]int{}
	for _, word := range words {
		result = append(result, map[string]int{word: 1})
	}
	return result
}

// Reduce function: Combines multiple word-count pairs into a single map
func reducer(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)
	for key, value := range m1 {
		result[key] = value + m2[key]
	}
	for key, value := range m2 {
		if _, ok := result[key]; !ok {
			result[key] = value
		}
	}
	return result
}

// Main function: Reads the file, applies Map and Reduce, and prints the result
func main() {
	data, err := ioutil.ReadFile("sample.txt") // Replace 'sample.txt' with your large text file
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	// Apply Map
	mappedResults := []map[string]int{}
	for _, line := range lines {
		mappedResults = append(mappedResults, mapper(line)...)
	}

	// Apply Reduce
	finalResult := mappedResults[0]
	for _, result := range mappedResults[1:] {
		finalResult = reducer(finalResult, result)
	}

	// Print the result
	for word, count := range finalResult {
		fmt.Printf("%s: %d\n", word, count)
	}
}

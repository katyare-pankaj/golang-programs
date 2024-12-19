package main

import (
	"fmt"
	"strconv"
)

// filterAndConvertStringsToInts filters out invalid strings and converts valid strings to integers.
func filterAndConvertStringsToInts(inputSlice []string) []int {
	// Create a new slice to store valid integers
	resultSlice := make([]int, 0, len(inputSlice))

	// Use a loop to filter and convert strings to integers
	for _, input := range inputSlice {
		if num, err := strconv.Atoi(input); err == nil {
			resultSlice = append(resultSlice, num)
		} else {
			fmt.Printf("Skipping invalid input: %q (error: %v)\n", input, err)
		}
	}

	return resultSlice
}

func main() {
	inputSlice := []string{"1", "two", "3.5", "four", "5", "nan", "7"}
	validIntegers := filterAndConvertStringsToInts(inputSlice)

	fmt.Println("Valid integers:", validIntegers)
}

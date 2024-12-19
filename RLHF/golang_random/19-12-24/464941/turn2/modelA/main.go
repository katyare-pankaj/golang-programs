package main

import (
	"fmt"
	"strconv"
)

// ConvertStringsToInts converts a slice of strings to a slice of integers,
// handling invalid entries by skipping them and returning the valid integers.
func ConvertStringsToInts(inputSlice []string) []int {
	resultSlice := make([]int, 0, len(inputSlice))

	for _, input := range inputSlice {
		num, err := strconv.Atoi(input)
		if err == nil {
			resultSlice = append(resultSlice, num)
		} else {
			fmt.Printf("Skipping invalid input: %q (error: %v)\n", input, err)
		}
	}

	return resultSlice
}

func main() {
	inputSlice := []string{"1", "2", "three", "4", "five", "6"}
	validIntegers := ConvertStringsToInts(inputSlice)

	fmt.Println("Valid integers:", validIntegers)
}

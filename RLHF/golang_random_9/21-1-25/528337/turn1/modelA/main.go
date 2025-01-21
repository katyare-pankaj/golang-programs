package main

import (
	"fmt"
)

// ImmutableSum calculates the sum of elements in a slice without modifying it.
func ImmutableSum(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}

// ImmutableFilter creates a new slice with elements that satisfy the predicate.
func ImmutableFilter(data []int, predicate func(int) bool) []int {
	var result []int
	for _, value := range data {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// ImmutableMap creates a new slice by applying a function to each element.
func ImmutableMap(data []int, transform func(int) int) []int {
	result := make([]int, len(data))
	for i, value := range data {
		result[i] = transform(value)
	}
	return result
}

func main() {
	// Original data slice
	data := []int{1, 2, 3, 4, 5}

	// Calculate the sum of the data
	sum := ImmutableSum(data)
	fmt.Printf("Sum: %d\n", sum)

	// Filter data to keep only even numbers
	evenData := ImmutableFilter(data, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("Even numbers: %v\n", evenData)

	// Map data to square each element
	squaredData := ImmutableMap(data, func(x int) int {
		return x * x
	})
	fmt.Printf("Squared numbers: %v\n", squaredData)

	// The original data is unchanged
	fmt.Printf("Original data: %v\n", data)
}

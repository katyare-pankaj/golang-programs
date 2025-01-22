package main

import (
	"fmt"
	"testing"
)

// SquareSlice squares each element of the input slice and returns a new slice with the squared values.
func SquareSlice(input []int) []int {
	// Initialize output slice with the same length as the input slice
	output := make([]int, len(input))

	// Iterate over input slice to square each element
	for i, val := range input {
		output[i] = val * val
	}
	return output
}

func TestSquareSlice(t *testing.T) {
	// Test case 1: Normal data
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 4, 9, 16, 25}
	result := SquareSlice(input)
	if !equal(result, expected) {
		t.Errorf("Test failed for input %v. Expected %v, got %v", input, expected, result)
	}

	// Test case 2: Empty slice
	input = []int{}
	expected = []int{}
	result = SquareSlice(input)
	if !equal(result, expected) {
		t.Errorf("Test failed for empty input. Expected %v, got %v", expected, result)
	}

	// Test case 3: Single element slice
	input = []int{8}
	expected = []int{64}
	result = SquareSlice(input)
	if !equal(result, expected) {
		t.Errorf("Test failed for single element input. Expected %v, got %v", expected, result)
	}
}

// Helper function to compare two slices for equality
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// Sample input slice
	inputData := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", inputData)

	// Square the elements of the input slice
	squaredData := SquareSlice(inputData)
	fmt.Println("Squared slice:", squaredData)

	// Run tests
	fmt.Println("Running tests...")
	t := &testing.T{}
	TestSquareSlice(t)
	if !t.Failed() {
		fmt.Println("All tests passed.")
	} else {
		fmt.Println("Some tests failed.")
	}
}

package main

import (
	"fmt"
	"testing"
)

func ETLProcess(inputData []int) []int {
	// Initialize the output slice with the same length as the input
	outputData := make([]int, len(inputData))

	// Use the 'range' keyword to iterate over the input slice and perform the ETL transformation
	for index, value := range inputData {
		// Replace each value with the result of the transformation (in this case, doubling the value)
		outputData[index] = value * 2
	}

	return outputData
}

func TestETLProcess(t *testing.T) {
	// Test case 1: Check if the ETL process correctly transforms the given input data
	testInput := []int{1, 2, 3, 4, 5}
	expectedOutput := []int{2, 4, 6, 8, 10}
	result := ETLProcess(testInput)
	if !equal(result, expectedOutput) {
		t.Errorf("ETL process failed for input %v. Expected %v, got %v", testInput, expectedOutput, result)
	}

	// Test case 2: Check if the ETL process handles an empty input slice
	testInput = []int{}
	expectedOutput = []int{}
	result = ETLProcess(testInput)
	if !equal(result, expectedOutput) {
		t.Errorf("ETL process failed for empty input. Expected %v, got %v", expectedOutput, result)
	}

	// Test case 3: Check if the ETL process handles a single-element input slice
	testInput = []int{10}
	expectedOutput = []int{20}
	result = ETLProcess(testInput)
	if !equal(result, expectedOutput) {
		t.Errorf("ETL process failed for single-element input. Expected %v, got %v", expectedOutput, result)
	}
}

func equal(a, b []int) bool {
	// Check if both slices have the same length
	if len(a) != len(b) {
		return false
	}

	// Compare each element of the slices
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	// Sample input data for ETL
	inputData := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Data: ", inputData)

	// Perform the ETL process
	transformedData := ETLProcess(inputData)
	fmt.Println("Transformed Data: ", transformedData)
}

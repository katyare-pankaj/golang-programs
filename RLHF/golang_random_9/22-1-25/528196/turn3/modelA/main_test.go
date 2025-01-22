package main

import (
	"fmt"
	"testing"
)

// squareSliceElements squares each element in the input slice.
// It handles negative integers by squaring them, which results in non-negative numbers.
func squareSliceElements(input []int) []int {
	result := make([]int, len(input))
	for i, v := range input {
		// Squaring both positive and negative integers results in a non-negative number
		result[i] = v * v
	}
	return result
}

func TestSquareSliceElements(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, -2, 3, -4, 5}, []int{1, 4, 9, 16, 25}},
		{[]int{}, []int{}},
		{[]int{-10}, []int{100}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test with input %v", testCase.input), func(t *testing.T) {
			result := squareSliceElements(testCase.input)
			if !equal(result, testCase.expected) {
				t.Errorf("Expected %v, got %v", testCase.expected, result)
			}
		})
	}
}

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
	testInput := []int{1, -2, 3, -4, 5}
	fmt.Println("Original Data: ", testInput)
	result := squareSliceElements(testInput)
	fmt.Println("Squared Data: ", result)
}

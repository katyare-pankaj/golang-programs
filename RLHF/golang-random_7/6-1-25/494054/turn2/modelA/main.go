package main

import (
	"fmt"
)

// Transformation represents a function that transforms an integer
type Transformation func(int) int

// ApplyTransformations applies a list of transformations sequentially to a list of integers
func ApplyTransformations(numbers []int, transformations []Transformation) []int {
	transformed := make([]int, len(numbers))
	for i, num := range numbers {
		for _, transform := range transformations {
			num = transform(num)
		}
		transformed[i] = num
	}
	return transformed
}

// Example transformations
func AddOne(x int) int {
	return x + 1
}

func MultiplyByTwo(x int) int {
	return x * 2
}

func EvenOnly(x int) int {
	if x%2 == 0 {
		return x
	}
	return 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Define the sequence of transformations
	transformations := []Transformation{
		AddOne,        // Add 1
		MultiplyByTwo, // Multiply by 2
		EvenOnly,      // Keep only even numbers
	}

	// Apply the transformations
	transformed := ApplyTransformations(numbers, transformations)

	// Print the transformed list
	fmt.Println("Original List:", numbers)
	fmt.Println("Transformed List:", transformed)
}

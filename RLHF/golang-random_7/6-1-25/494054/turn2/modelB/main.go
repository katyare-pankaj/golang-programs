package main

import (
	"fmt"
)

// TransformFunction is a function that takes an integer as input and returns an integer.
type TransformFunction func(int) int

// ApplyTransformPipeline applies a sequence of transformations to a list of integers.
func ApplyTransformPipeline(list []int, transformations []TransformFunction) []int {
	result := list
	for _, transform := range transformations {
		transformedResult := make([]int, len(result))
		for i, val := range result {
			transformedResult[i] = transform(val)
		}
		result = transformedResult
	}
	return result
}

func double(x int) int {
	return x * 2
}

func increment(x int) int {
	return x + 1
}

func subtractFive(x int) int {
	return x - 5
}

func main() {
	numList := []int{1, 2, 3, 4, 5}

	// Define a pipeline of transformations
	pipeline := []TransformFunction{
		double,
		increment,
		subtractFive,
	}

	fmt.Println("Original list: ", numList)
	fmt.Println("Transformed list: ", ApplyTransformPipeline(numList, pipeline))
}

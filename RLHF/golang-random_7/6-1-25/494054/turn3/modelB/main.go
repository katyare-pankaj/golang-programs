package main

import (
	"errors"
	"fmt"
)

// TransformFunction is a function that takes an integer as input and returns an integer,
// and also returns an error if it occurs.
type TransformFunction func(int) (int, error)

// ApplyTransformPipeline applies a sequence of transformations to a list of integers.
func ApplyTransformPipeline(list []int, transformations []TransformFunction) ([]int, error) {
	result := list
	for _, transform := range transformations {
		transformedResult := make([]int, len(result))
		for i, val := range result {
			transformedVal, err := transform(val)
			if err != nil {
				return nil, err
			}
			transformedResult[i] = transformedVal
		}
		result = transformedResult
	}
	return result, nil
}

func double(x int) (int, error) {
	return x * 2, nil
}

func increment(x int) (int, error) {
	return x + 1, nil
}

func divideByZero(x int) (int, error) {
	return 0, errors.New("division by zero")
}

func main() {
	numList := []int{1, 2, 3, 4, 5}

	// Define a pipeline of transformations
	pipeline := []TransformFunction{
		double,
		increment,
		divideByZero,
	}

	transformedList, err := ApplyTransformPipeline(numList, pipeline)
	if err != nil {
		fmt.Println("Error occurred during transformation:", err)
	} else {
		fmt.Println("Transformed list: ", transformedList)
	}
}

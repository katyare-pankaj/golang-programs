package main

import (
	"errors"
	"fmt"
)

// TransformFunction is a function that takes an integer as input and returns an integer and an error.
type TransformFunction func(int) (int, error)

// ApplyTransformPipeline applies a sequence of transformations to a list of integers, handling errors.
func ApplyTransformPipeline(list []int, transformations []TransformFunction) ([]int, []error) {
	result := list
	errorsEncountered := make([]error, len(result))
	for i, val := range result {
		var err error
		result[i], err = applyTransformations(val, transformations)
		errorsEncountered[i] = err
	}
	return result, errorsEncountered
}

func applyTransformations(val int, transformations []TransformFunction) (int, error) {
	currentValue := val
	for _, transform := range transformations {
		var err error
		currentValue, err = transform(currentValue)
		if err != nil {
			return currentValue, err
		}
	}
	return currentValue, nil
}

func double(x int) (int, error) {
	return x * 2, nil
}

func increment(x int) (int, error) {
	return x + 1, nil
}

func divideByZero(x int) (int, error) {
	if x == 0 {
		return 0, errors.New("division by zero")
	}
	return x, nil
}

func main() {
	numList := []int{1, 2, 3, 0, 5}

	// Define a pipeline of transformations
	pipeline := []TransformFunction{
		double,
		increment,
		divideByZero,
	}

	fmt.Println("Original list: ", numList)
	transformedList, errorsEncountered := ApplyTransformPipeline(numList, pipeline)

	fmt.Println("Transformed list: ", transformedList)
	for _, err := range errorsEncountered {
		if err != nil {
			fmt.Println("Error encountered:", err)
		}
	}
}

package main

import (
	"errors"
	"fmt"
)

// CustomError is a type that includes additional context about an error
type CustomError struct {
	Function string
	Err      error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error in %s: %v", e.Function, e.Err)
}

// A function that performs a division and returns an error if the divisor is zero
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &CustomError{Function: "divide", Err: errors.New("division by zero")}
	}
	return a / b, nil
}

// A function that calculates an average and propagates any errors from divide function
func calculateAverage(sum float64, count float64) (float64, error) {
	average, err := divide(sum, count)
	if err != nil {
		return 0, &CustomError{Function: "calculateAverage", Err: err}
	}
	return average, nil
}

// A function that computes the result of some operation (e.g., average calculation)
func computeResult(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, &CustomError{Function: "computeResult", Err: errors.New("empty slice")}
	}

	var sum float64
	for _, number := range numbers {
		sum += number
	}

	result, err := calculateAverage(sum, float64(len(numbers)))
	if err != nil {
		return 0, &CustomError{Function: "computeResult", Err: err}
	}
	return result, nil
}

func main() {
	// Test cases
	tests := [][]float64{
		{},           // Empty slice
		{1, 2, 3, 4}, // Normal input
		{10, 20, 30}, // Another normal input
	}

	// Execute test cases
	for i, test := range tests {
		fmt.Printf("Running test case %d: %v\n", i+1, test)
		result, err := computeResult(test)
		if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
		} else {
			fmt.Printf("The result is: %f\n", result)
		}
		fmt.Println()
	}
}

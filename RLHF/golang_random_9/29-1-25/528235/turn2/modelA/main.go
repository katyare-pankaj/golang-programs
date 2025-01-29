package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Simulated resource-intensive function
func ResourceIntensiveOp(param int) (string, error) {
	fmt.Printf("Processing with param: %d\n", param)
	time.Sleep(500 * time.Millisecond) // Simulating heavy operation

	if param%2 == 0 { // Simulate an error condition
		return "", errors.New("even number error")
	}
	return fmt.Sprintf("Success with %d", param), nil
}

// Function to invoke with retries
func InvokeWithRetries(function interface{}, params ...interface{}) (interface{}, error) {
	const maxRetries = 3

	// Convert function and parameters into reflect.Values
	fnValue := reflect.ValueOf(function)
	inParams := make([]reflect.Value, len(params))
	for i, param := range params {
		inParams[i] = reflect.ValueOf(param)
	}

	// Retry mechanism
	var result []reflect.Value
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		result = fnValue.Call(inParams)

		// Check for error return
		if err, _ = result[1].Interface().(error); err != nil {
			fmt.Printf("Attempt %d failed: %v\n", attempt, err)
			time.Sleep(1 * time.Second) // Retry delay
			continue
		}

		// Successful call
		return result[0].Interface(), nil
	}

	return nil, fmt.Errorf("operation failed after %d retries", maxRetries)
}

func main() {
	// Resource cleanup with defer
	defer finalizeResources()

	// Try invoking with a parameter that will cause retries
	if res, err := InvokeWithRetries(ResourceIntensiveOp, 4); err != nil {
		fmt.Println("Final failure:", err)
	} else {
		fmt.Println("Operation result:", res)
	}

	// Try invoking with a parameter that will succeed
	if res, err := InvokeWithRetries(ResourceIntensiveOp, 3); err != nil {
		fmt.Println("Final failure:", err)
	} else {
		fmt.Println("Operation result:", res)
	}
}

// Simulated resource finalization function
func finalizeResources() {
	fmt.Println("Cleaning up resources...")
}

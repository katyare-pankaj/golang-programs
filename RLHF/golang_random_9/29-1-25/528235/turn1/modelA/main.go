package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Example of a resource-intensive function
func ResourceIntensiveOperation(param1 int, param2 string) (string, error) {
	// Simulate some processing
	time.Sleep(1 * time.Second)

	// Simulate an error condition to trigger retry
	if param1%2 == 0 {
		return "", errors.New("simulated error")
	}

	return fmt.Sprintf("Processed: %d, %s", param1, param2), nil
}

// Function to dynamically invoke using reflection
func InvokeWithRetry(fn interface{}, params ...interface{}) (interface{}, error) {
	const maxRetries = 3
	var result []reflect.Value

	fnValue := reflect.ValueOf(fn)
	inParams := make([]reflect.Value, len(params))
	for i, param := range params {
		inParams[i] = reflect.ValueOf(param)
	}

	// Retry mechanism
	for try := 0; try < maxRetries; try++ {
		result = fnValue.Call(inParams)
		if err, ok := result[1].Interface().(error); ok && err != nil {
			fmt.Printf("Attempt %d failed: %v\n", try+1, err)
			time.Sleep(500 * time.Millisecond) // retry delay
			continue
		}
		// Successful call
		return result[0].Interface(), nil
	}
	return nil, errors.New("all retries failed")
}

func main() {
	// Ensure proper resource management with defer
	defer finalizeResources()

	res, err := InvokeWithRetry(ResourceIntensiveOperation, 4, "example")
	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
		return
	}
	fmt.Printf("Operation succeeded: %v\n", res)
}

// Simulated function to finalize resources
func finalizeResources() {
	fmt.Println("Finalizing resources...")
}

package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type Resource struct {
	// Add any necessary resource-specific fields
}

func (r *Resource) Init() error {
	// Simulate resource initialization
	fmt.Println("Initializing resource...")
	time.Sleep(1 * time.Second)
	return nil
}

func (r *Resource) Close() {
	// Simulate resource cleanup
	fmt.Println("Cleaning up resource...")
}

func (r *Resource) ResourceIntensiveOperation(param1 int, param2 string) (string, error) {
	// Simulate resource-intensive operation
	time.Sleep(1 * time.Second)

	// Simulate an error condition to trigger retry
	if param1%2 == 0 {
		return "", errors.New("simulated error")
	}

	return fmt.Sprintf("Processed: %d, %s", param1, param2), nil
}

// Function to dynamically invoke using reflection with retry and resource management
func InvokeWithRetryAndResource(fn interface{}, resource interface{}, params ...interface{}) (interface{}, error) {
	const maxRetries = 3
	var result []reflect.Value

	// Initialize resource using reflection
	resValue := reflect.ValueOf(resource)
	initMethod := resValue.MethodByName("Init")
	if !initMethod.IsValid() {
		return nil, errors.New("resource does not have 'Init' method")
	}
	initErr := initMethod.Call(nil)[0].Interface().(error)
	if initErr != nil {
		return nil, initErr
	}
	defer func() {
		// Defer the resource cleanup using reflection
		closeMethod := resValue.MethodByName("Close")
		if !closeMethod.IsValid() {
			fmt.Println("Warning: Resource does not have 'Close' method")
			return
		}
		closeMethod.Call(nil)
	}()

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
	var r Resource
	res, err := InvokeWithRetryAndResource(r.ResourceIntensiveOperation, &r, 4, "example")
	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
		return
	}
	fmt.Printf("Operation succeeded: %v\n", res)
}

package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Example function to be invoked dynamically
func DoWork(param int) (string, error) {
	if param < 0 {
		return "", errors.New("param cannot be negative")
	}
	return fmt.Sprintf("Success with %d!", param), nil
}

// Invoke any function with error handling using reflection
func InvokeFunction(fn interface{}, args ...interface{}) (interface{}, error) {
	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		return nil, errors.New("provided argument is not a function")
	}

	// Prepare input arguments
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// Perform the function call
	defer func(start time.Time) {
		// Addition of some cleanup logic or resource management, if needed
		elapsed := time.Since(start)
		fmt.Printf("Function execution took %s\n", elapsed)
	}(time.Now())

	results := fnValue.Call(in)

	// Assume the last return value is an error, which is common practice in Go
	if len(results) > 1 && !results[len(results)-1].IsNil() {
		err := results[len(results)-1].Interface().(error)
		return nil, err
	}

	return results[0].Interface(), nil
}

func main() {
	// Ensure any necessary cleanup using defer
	defer fmt.Println("Main execution completed.")

	// Example usage
	if result, err := InvokeFunction(DoWork, -1); err != nil {
		fmt.Println("Error invoking function:", err)
	} else {
		fmt.Println("Function result:", result)
	}

	if result, err := InvokeFunction(DoWork, 10); err != nil {
		fmt.Println("Error invoking function:", err)
	} else {
		fmt.Println("Function result:", result)
	}
}

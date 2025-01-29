package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"golang.org/x/exp/rand"
)

type MyService struct {
	// Add any resource-intensive variables or methods here.
}

// Example resource-intensive method.
func (m *MyService) ResourceIntensiveMethod() error {
	// Simulate resource-intensive work by sleeping for a random duration.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Some random error scenario.
	if rand.Intn(10) == 0 {
		return errors.New("random error")
	}

	fmt.Println("Resource-intensive method executed successfully.")
	return nil
}

func retryMethod(m interface{}, methodName string, numRetries int, retryDelay time.Duration) (err error) {
	v := reflect.ValueOf(m)
	method := v.MethodByName(methodName)
	if !method.IsValid() {
		return errors.New("method not found")
	}

	for i := 0; i <= numRetries; i++ {
		if i > 0 {
			fmt.Println("Retrying...")
		}

		results := method.Call(nil) // Invoke the method with no arguments.
		if len(results) > 0 && results[0].Interface() != nil {
			err = results[0].Interface().(error)
			if err == nil {
				return nil // Success, return without error.
			}
			fmt.Println("Error:", err)
		}

		// Retry after delay.
		time.Sleep(retryDelay)
	}

	return fmt.Errorf("retries exhausted, last error: %w", err)
}

func main() {
	service := &MyService{}

	// Example usage:
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	err := retryMethod(service, "ResourceIntensiveMethod", 3, time.Second) // Retry up to 3 times with a 1-second delay.
	if err != nil {
		fmt.Println("Final Error:", err)
	}
}

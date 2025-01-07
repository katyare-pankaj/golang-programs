package main

import (
	"fmt"
)

// Custom error struct
type CustomError struct {
	message string
	code    int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.code, e.message)
}

// Function returning multiple values with a custom error
func processData() (data string, size int, err error) {
	// Simulated error
	if true {
		return "", 0, CustomError{message: "Data processing error", code: 400}
	}
	// Simulated valid result
	return "Processed data", 2048, nil
}

func main() {
	// Calling the function
	data, size, err := processData()
	if err != nil {
		if cerr, ok := err.(CustomError); ok {
			fmt.Printf("Custom error: %v\n", cerr)
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
		return
	}

	fmt.Printf("Data: %s, Size: %d\n", data, size)
}

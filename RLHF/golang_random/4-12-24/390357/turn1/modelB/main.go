package main

import (
	"fmt"
	"log"
)

// SafeDivide performs division and recovers from any panics.
func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic
			log.Printf("Recovered from panic: %v", r)
			err = fmt.Errorf("division failed: %v", r)
		}
	}()

	// Check for division by zero
	if b == 0 {
		panic("division by zero")
	}

	return a / b, nil
}

func main() {
	// Example usage
	if result, err := SafeDivide(10, 2); err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	// This will cause a panic, which will be recovered
	if result, err := SafeDivide(10, 0); err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}

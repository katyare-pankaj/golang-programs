package main

import (
	"errors"
	"fmt"
)

// Define custom errors using errors.New
var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrDivisionByZero = errors.New("division by zero")
)

// Use error chaining for more context
func divide(a, b int) error {
	if b == 0 {
		return fmt.Errorf("%w: cannot divide by zero", ErrDivisionByZero)
	}
	return nil
}

func processData(input []int) error {
	for _, value := range input {
		if value < 0 {
			return fmt.Errorf("%w: invalid negative value: %d", ErrInvalidInput, value)
		}
		err := divide(value, 2)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	input := []int{1, 2, -3, 4, 0}
	if err := processData(input); err != nil {
		switch err {
		case ErrInvalidInput:
			fmt.Println("Invalid input error:", err)
		case ErrDivisionByZero:
			fmt.Println("Division by zero error:", err)
		default:
			fmt.Println("Unexpected error:", err)
		}
	} else {
		fmt.Println("Processing completed successfully")
	}
}

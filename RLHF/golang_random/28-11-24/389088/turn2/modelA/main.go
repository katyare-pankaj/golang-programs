package main

import (
	"fmt"
	"time"
)

type ErrInvalidInput struct {
	message string
}

func (e *ErrInvalidInput) Error() string {
	return e.message
}

func NewErrInvalidInput(message string) error {
	return &ErrInvalidInput{message: message}
}

type ErrInternalError struct {
	message string
}

func (e *ErrInternalError) Error() string {
	return e.message
}

func NewErrInternalError(message string) error {
	return &ErrInternalError{message: message}
}

func validateInput(input string) error {
	if len(input) == 0 {
		return NewErrInvalidInput("Input cannot be empty")
	}
	return nil
}

func processInput(input string) error {
	// Simulate some processing
	time.Sleep(100 * time.Millisecond)
	return nil
}

func main() {
	userInput := "hello" // Example user input

	err := validateInput(userInput)
	if err != nil {
		switch err.(type) {
		case *ErrInvalidInput:
			fmt.Println("Error:", err.Error()) // User-friendly message for user
		default:
			fmt.Println("An unexpected error occurred while validating input:", err) // Internal logging for developer
		}
		return
	}

	err = processInput(userInput)
	if err != nil {
		switch err.(type) {
		case *ErrInternalError:
			fmt.Println("An internal error occurred:", err) // Informative message for user
		default:
			fmt.Println("A critical error occurred:", err) // General error message for user
			fmt.Println("Details:", err)                   // Internal logging for developer
		}
		return
	}

	fmt.Println("Input processed successfully.")
}

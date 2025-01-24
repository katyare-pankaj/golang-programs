package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// CustomError is a struct that includes additional context about an error
type CustomError struct {
	Function string
	Err      error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error in %s: %v", e.Function, e.Err)
}

// readInput reads input from the console and returns an error if the input is empty
func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", &CustomError{Function: "readInput", Err: err}
	}

	trimmedInput := strings.TrimSpace(input)
	if trimmedInput == "" {
		return "", &CustomError{Function: "readInput", Err: errors.New("input cannot be empty")}
	}

	return trimmedInput, nil
}

// processInput converts the input to uppercase
func processInput(input string) (string, error) {
	if input == "" {
		return "", &CustomError{Function: "processInput", Err: errors.New("empty input provided to processing")}
	}
	return strings.ToUpper(input), nil
}

// execute is the top-level function that chains the input reading and processing
func execute() (string, error) {
	input, err := readInput()
	if err != nil {
		return "", &CustomError{Function: "execute", Err: err}
	}

	processedInput, err := processInput(input)
	if err != nil {
		return "", &CustomError{Function: "execute", Err: err}
	}

	return processedInput, nil
}

func main() {
	result, err := execute()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	} else {
		fmt.Printf("Processed input: %s\n", result)
	}
}

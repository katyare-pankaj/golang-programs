package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Function to read input from user
func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your input: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", &CustomError{Function: "readInput", Err: err}
	}
	text = strings.TrimSpace(text)
	return text, nil
}

// Function to convert input to uppercase
func convertToUppercase(input string) (string, error) {
	if input == "" {
		return "", &CustomError{Function: "convertToUppercase", Err: errors.New("input is empty")}
	}
	return strings.ToUpper(input), nil
}

// Custom error struct to track error origin
type CustomError struct {
	Function string
	Err      error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error in %s: %v", e.Function, e.Err)
}

// Main function
func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println("Error occurred while reading input:", err)
		return
	}

	result, err := convertToUppercase(input)
	if err != nil {
		fmt.Println("Error occurred while converting input:", err)
		return
	}
	fmt.Println("Converted input to uppercase:", result)
}

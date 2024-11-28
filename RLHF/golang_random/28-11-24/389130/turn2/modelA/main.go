package main

import (
	"fmt"
	"os"
)

// Define a custom error type
type FileError struct {
	path string
	err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("error opening file '%s': %w", e.path, e.err)
}

func (e *FileError) Unwrap() error {
	return e.err
}

// Function to read a file
func readFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		// Wrap the original error with context
		return &FileError{path: filePath, err: err}
	}
	defer file.Close()
	return nil
}

func main() {
	filePath := "non-existent-file.txt"

	err := readFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		// Access the original error if needed
		if fileErr, ok := err.(*FileError); ok {
			fmt.Printf("Original error: %v\n", fileErr.Unwrap())
		}
	} else {
		fmt.Println("File read successfully.")
	}
}

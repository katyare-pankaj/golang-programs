package main

import (
	"fmt"
	"os"
)

func main() {
	if err := processFile("sample.txt"); err != nil {
		fmt.Println("Error processing file:", err)
		os.Exit(1)
	}
	fmt.Println("File processed successfully")
}

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	// Ensure this file is closed when the function returns
	defer func() {
		fmt.Println("Closing the file")
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Simulating file processing
	if err := readData(file); err != nil {
		return fmt.Errorf("failed to read from file: %w", err)
	}

	// Simulate error-free completion
	return nil
}

func readData(file *os.File) error {
	// In a realistic scenario this would be reading and processing data
	// Here, simulate an error condition
	return fmt.Errorf("mock read error")
}

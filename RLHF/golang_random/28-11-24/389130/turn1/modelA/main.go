package main

import (
	"fmt"
	"os"
)

func readFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		// Create a structured error with context
		return fmt.Errorf("error opening file '%s': %w", filePath, err)
	}
	defer file.Close()
	// Continue with file operations
	return nil
}

func main() {
	filePath := "D:\\Companygo\\golang-programs\\RLHF\\golang_random\\28-11-24\\389130\\turn1\\modelA\\example.txt"
	err := readFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		// Optionally, you can log additional details for debugging
		fmt.Println("Debug info: Check file path and permissions.")
	} else {
		fmt.Println("File read successfully.")
	}
}

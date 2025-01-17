package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if err := processFile("example.txt"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File processed successfully.")
	}
}

func processFile(filename string) error {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	// Ensure the file is closed at the end of the function
	defer file.Close()

	// Simulate file processing that might fail
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Do something with the file content
	fmt.Printf("File content: %s\n", content)

	// Simulate another potential failure
	if len(content) == 0 {
		return fmt.Errorf("file is empty")
	}

	return nil
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// readFile opens a file and reads its contents, ensuring the file is closed afterwards
func readFile(filename string) error {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	// Defer the file close operation
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print each line
	}

	// Check for scanning errors
	if scanErr := scanner.Err(); scanErr != nil {
		return fmt.Errorf("error reading file: %v", scanErr)
	}

	return nil
}

// writeFile opens a file (creating/truncating it) and writes some content to it
func writeFile(filename string, content []string) error {
	// Open the file for writing with create/truncate flags
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	// Defer the file close operation
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	// Create a new writer for the file
	writer := bufio.NewWriter(file)
	for _, line := range content {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	// Ensure all buffered operations are applied
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing buffer: %v", err)
	}

	return nil
}

func main() {
	// Example content to write
	content := []string{
		"Hello, World!",
		"This is a test file.",
		"Using defer in Go.",
	}

	// Write example content to file
	if err := writeFile("example.txt", content); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	// Read from the same file
	if err := readFile("example.txt"); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}

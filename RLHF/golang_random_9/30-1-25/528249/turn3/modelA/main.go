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

// writeFile opens a file with append flag and writes content to it
func appendToFile(filename string, content []string) error {
	// Open the file for appending with appropriate flags
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	// Create a buffered writer for the file
	writer := bufio.NewWriter(file)
	for _, line := range content {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	// Ensure all buffered data is written to the file
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing buffer: %v", err)
	}

	return nil
}

func main() {
	// Example content to append
	content := []string{
		"Appending line 1.",
		"Appending line 2.",
		"Appending line 3.",
	}

	// Append example content to file
	if err := appendToFile("example.txt", content); err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	// Read from the file to verify append
	if err := readFile("example.txt"); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}

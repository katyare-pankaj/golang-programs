package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ... (Previous code remains the same)

// writeFile opens a file (appending to it if it exists) and writes some content to it.
func writeFile(filename string, content []string) error {
	// Open the file for writing (appending) with create flag if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	// Existing content in the file
	existingContent := []string{"This line will be kept.", "The next line will also be kept."}
	// New content to append
	contentToAppend := []string{"Appending this line.", "Adding another line."}

	// Write existing content to the file
	if err := writeFile("example.txt", existingContent); err != nil {
		log.Fatalf("Error writing existing content to file: %v", err)
	}

	// Write the content to be appended to the file
	if err := writeFile("example.txt", contentToAppend); err != nil {
		log.Fatalf("Error appending content to file: %v", err)
	}

	// Rest of the code remains the same...
}

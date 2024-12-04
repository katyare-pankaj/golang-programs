package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}

	// Use defer to ensure file is closed regardless of how the function exits
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Process each line (for demonstration, print it)
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %v", err)
	}
}

func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}

	for _, filename := range filenames {
		// Start a new goroutine for each file
		go processFile(filename)
	}

	// Wait for all goroutines to finish
	fmt.Println("Waiting for goroutines to complete...")
	for _, filename := range filenames {
		// You can use channels or other synchronization methods if needed
	}
	fmt.Println("All goroutines completed.")
}

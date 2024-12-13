package main

import (
	"fmt"
	"os"
	"sync"
)

// Function that simulates processing data from a file
func processDataConcurrently(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Simulate reading and processing data from the file
	fmt.Println("Processing data from", fileName)

	// In a real scenario, you would read from the file here
}

// Main function
func main() {
	var wg sync.WaitGroup
	fileNames := []string{"file1.txt", "file2.txt", "file3.txt"}

	// Start goroutines to process each file concurrently
	for _, fileName := range fileNames {
		wg.Add(1)
		go processDataConcurrently(fileName, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All files have been processed.")
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// Function to read a file and return its content
func readFile(filePath string, wg *sync.WaitGroup, results chan<- string, errors chan<- error) {
	defer wg.Done() // Signal that this Goroutine is done

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		errors <- err // Send error to the errors channel
		return
	}
	defer file.Close() // Ensure file is closed when we're done

	// Read the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		errors <- err // Send error to the errors channel
		return
	}

	results <- string(content) // Send result to the results channel
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string)
	errors := make(chan error)

	filePaths := []string{"file1.txt", "file2.txt", "file3.txt"}

	// Start Goroutines
	for _, path := range filePaths {
		wg.Add(1)
		go readFile(path, &wg, results, errors)
	}

	// Close results and errors channels once all Goroutines are done
	go func() {
		wg.Wait()      // Wait for all Goroutines to finish
		close(results) // Close the results channel
		close(errors)  // Close the errors channel
	}()

	// Collect results and errors
	for {
		select {
		case result, ok := <-results:
			if ok {
				fmt.Println("File content:", result)
			}
		case err, ok := <-errors:
			if ok {
				log.Println("Error:", err)
			}
		case <-results: // Break out when results channel is closed
			return
		}
	}
}

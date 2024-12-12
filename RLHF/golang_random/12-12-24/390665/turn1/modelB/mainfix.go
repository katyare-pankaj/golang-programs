package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func readFileConcurrently(ctx context.Context, filePaths []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, filePath := range filePaths {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine for %s cancelled\n", filePath)
			return
		default:
			// Read the file (but we don't need to store the data in a variable)
			_, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", filePath, err)
				continue
			}

			// Simulate processing time
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

			fmt.Printf("Processed %s\n", filePath)
		}
	}
}

func chunkFiles(filePaths []string, numChunks int) [][]string {
	var chunks [][]string
	chunkSize := (len(filePaths) + numChunks - 1) / numChunks

	for i := 0; i < len(filePaths); i += chunkSize {
		end := i + chunkSize
		if end > len(filePaths) {
			end = len(filePaths)
		}
		chunks = append(chunks, filePaths[i:end])
	}
	return chunks
}

func main() {
	// List of file paths
	filePaths := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
		"file4.txt",
		"file5.txt",
	}

	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Number of goroutines
	numGoroutines := 3

	// Split the filePaths into chunks
	chunks := chunkFiles(filePaths, numGoroutines)

	// Start goroutines to read files concurrently
	for _, chunk := range chunks {
		wg.Add(1)
		go readFileConcurrently(ctx, chunk, &wg)
	}

	// Wait for all goroutines to finish or context to be cancelled
	wg.Wait()

	fmt.Println("All file operations completed or cancelled.")
}

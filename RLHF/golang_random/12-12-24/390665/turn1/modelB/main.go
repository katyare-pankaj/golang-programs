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
			// Read the file
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", filePath, err)
				continue
			}

			// Simulate processing time
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

			fmt.Printf("Read and processed %s\n", filePath)
		}
	}
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

	// Start goroutines to read files concurrently
	numGoroutines := 3
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go readFileConcurrently(ctx, filePaths, &wg)
	}

	// Wait for all goroutines to finish or context to be cancelled
	wg.Wait()

	fmt.Println("All file operations completed or cancelled.")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulated file download function
func downloadFile(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that the Goroutine is done
	fmt.Printf("Downloading file: %s\n", url)
	time.Sleep(time.Second * 2) // Simulate download duration
	fmt.Printf("File download completed: %s\n", url)
}

func main() {
	var wg sync.WaitGroup // Initialize WaitGroup
	urls := []string{
		"http://example.com/file1.txt",
		"http://example.com/file2.txt",
		"http://example.com/file3.txt",
		"http://example.com/file4.txt",
	}

	// Start Goroutines for each file download
	for _, url := range urls {
		wg.Add(1) // Increment the waitgroup for each download
		go downloadFile(url, &wg)
	}

	// Main thread waits for all downloads to complete
	fmt.Println("Main thread waiting for all downloads to finish...")
	wg.Wait()

	fmt.Println("All downloads completed. Main thread exiting.")
}

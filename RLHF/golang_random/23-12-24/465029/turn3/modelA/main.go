package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()             // Call Done() when the goroutine completes
	time.Sleep(2 * time.Second) // Simulate download time
	fmt.Printf("Downloaded file: %s\n", fileName)
}

func main() {
	var wg sync.WaitGroup

	filesToDownload := []string{"file1.txt", "file2.txt", "file3.txt"}

	// Increment the WaitGroup by the number of files we'll download
	wg.Add(len(filesToDownload))

	// Start goroutines to download files
	for _, file := range filesToDownload {
		go downloadFile(file, &wg)
	}

	// The main thread will wait here until all goroutines are done
	wg.Wait()

	// Execute cleanup or post-download tasks
	fmt.Println("All downloads complete. Performing cleanup...")
}

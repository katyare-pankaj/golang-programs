package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

// downloadFile downloads a file from the given URL and saves it to the local disk
func downloadFile(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	// Extract the file name from the URL
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Printf("Starting download: %s\n", fileName)

	// HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to download %s: %v\n", fileName, err)
		return
	}
	defer resp.Body.Close()

	// Create a file to save the download
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	// Write the body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Failed to save file %s: %v\n", fileName, err)
		return
	}

	fmt.Printf("Completed download: %s\n", fileName)
}

func main() {
	// List of URLs to download
	urls := []string{
		"https://example.com/file1.jpg",
		"https://example.com/file2.jpg",
		"https://example.com/file3.jpg",
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Add the number of URLs to process to the WaitGroup's counter
	wg.Add(len(urls))

	// Launch a goroutine to download each file
	for _, url := range urls {
		go downloadFile(url, &wg)
	}

	// Wait for all goroutines to finish downloading
	wg.Wait()
	fmt.Println("All downloads completed.")
}

package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Function to fetch the status code of a given URL
func fetchStatusCode(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", url, "->", err)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("Error closing response body for URL:", url, "->", err)
		}
	}()

	fmt.Println("Status code for URL:", url, "->", resp.StatusCode)
}

// Main function
func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://www.example.com",
		"https://go.dev",
		"https://github.com",
		"https://non-existent-url.com", // This will cause an error
	}

	// Start goroutines to fetch status codes concurrently
	for _, url := range urls {
		wg.Add(1)
		go fetchStatusCode(url, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All URLs have been processed.")
}

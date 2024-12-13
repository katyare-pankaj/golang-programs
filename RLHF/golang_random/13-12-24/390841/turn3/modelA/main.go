package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Function to fetch HTTP status code for a URL
func fetchURLStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error closing response body for %s: %v\n", url, err)
		}
	}()

	fmt.Printf("Status code for %s: %s\n", url, resp.Status)
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.nonexistenturl.xyz",
		"https://httpbin.org/status/404",
	}

	// Start goroutines to fetch status codes concurrently
	for _, url := range urls {
		wg.Add(1)
		go fetchURLStatus(url, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Finished fetching URLs.")
}

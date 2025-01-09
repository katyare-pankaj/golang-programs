package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Define a struct to represent a URL
type URL struct {
	URL string
}

// Function to crawl a URL and print its response code
func crawlURL(url URL, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Crawling URL: %s\n", url.URL)
	resp, err := http.Get(url.URL)
	if err != nil {
		fmt.Printf("Error crawling URL %s: %v\n", url.URL, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Response code for URL %s: %d\n", url.URL, resp.StatusCode)
}
func main() {
	// Create a bufferred channel to hold URLs to be crawled
	urlQueue := make(chan URL, 100)
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	// Start a fixed number of worker goroutines
	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for url := range urlQueue {
				crawlURL(url, &wg)
			}
			wg.Done()
		}()
	}
	// Send URLs to the channel for crawling
	urls := []string{
		"https://www.example.com",
		"https://www.golang.org",
		"https://www.youtube.com",
		// Add more URLs here
	}
	for _, url := range urls {
		urlQueue <- URL{URL: url}
	}
	// Close the channel to signal the end of input
	close(urlQueue)
	// Wait for all worker goroutines to finish
	wg.Wait()
	fmt.Println("All URLs crawled!")
}

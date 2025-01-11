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

// Define a channel to send URLs between goroutines
var urlChannel chan URL

// Initialize the channel with a buffer size of 100
func init() {
	urlChannel = make(chan URL, 100)
}

// Fetch the content of a URL
func fetchURL(url URL, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url.URL)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()
	// Process the fetched content (e.g., parse HTML, extract links)
	// For this example, we'll just print the status code
	fmt.Println("Fetched:", url.URL, "- Status:", resp.StatusCode)
}

// Function to add URLs to the queue
func addURLsToQueue() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.golang.org",
		"https://www.example.com",
		// Add more URLs as needed
	}
	for _, url := range urls {
		urlChannel <- URL{URL: url}
	}
	close(urlChannel)
}

// Main function
func main() {
	var wg sync.WaitGroup

	// Start the goroutine to add URLs to the queue
	go addURLsToQueue()

	// Start multiple goroutines to fetch URLs from the queue
	numGoroutines := 5 // Adjust this as needed
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			for url := range urlChannel {
				fetchURL(url, &wg)
			}
			wg.Done()
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All URLs fetched successfully.")
}

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// makeHTTPRequest simulates an HTTP request and updates the result channel
func makeHTTPRequest(wg *sync.WaitGroup, url string, results chan<- string) {
	defer wg.Done()
	fmt.Printf("Making HTTP request to: %s\n", url)

	// Simulate the HTTP request taking some time
	time.Sleep(time.Duration(2) * time.Second)

	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error for %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		results <- fmt.Sprintf("HTTP error for %s: %d", url, resp.StatusCode)
		return
	}

	results <- fmt.Sprintf("Success for %s: %d", url, resp.StatusCode)
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string)
	defer close(results)

	// List of URLs to make HTTP requests to
	urls := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
		"https://api.example.com/data4",
	}

	wg.Add(len(urls))

	// Start the HTTP requests in goroutines
	for _, url := range urls {
		go makeHTTPRequest(&wg, url, results)
	}

	// Wait for all requests to complete
	wg.Wait()

	fmt.Println("All HTTP requests completed. Results:")
	// Process and print results
	for result := range results {
		fmt.Println(result)
	}
}

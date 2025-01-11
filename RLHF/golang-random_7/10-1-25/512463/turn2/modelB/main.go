package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func makeAPICall(url string, wg *sync.WaitGroup, statusCodes chan<- int) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		// Handle error gracefully
		log.Printf("Error making API call to %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Log the HTTP status code
	statusCodes <- resp.StatusCode
}

func main() {
	urls := []string{
		"https://example.com",
		"https://www.example.com",
		"https://example.com/not-found",
		"https://www.example.com/page",
		"invalid-url",
		"https://example.com/api/data",
	}

	var wg sync.WaitGroup
	statusCodes := make(chan int, len(urls))

	wg.Add(len(urls))

	// Perform API calls concurrently
	for _, url := range urls {
		go makeAPICall(url, &wg, statusCodes)
	}

	go func() {
		wg.Wait()
		close(statusCodes)
	}()

	// Print the HTTP status codes for each URL
	for code := range statusCodes {
		fmt.Println("HTTP Status Code:", code)
	}

	fmt.Println("All API calls completed.")
}

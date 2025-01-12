package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	numRequests    = 10
	maxConcurrency = 3
)

func makeRequest(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request to %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	// Simulate processing the response
	time.Sleep(1 * time.Second)
	fmt.Printf("Request to %s completed with status: %d\n", url, resp.StatusCode)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://example.com/api/1",
		"https://example.com/api/2",
		"https://example.com/api/3",
		"https://example.com/api/4",
		"https://example.com/api/5",
		"https://example.com/api/6",
		"https://example.com/api/7",
		"https://example.com/api/8",
		"https://example.com/api/9",
		"https://example.com/api/10",
	}

	// Create a channel to limit the concurrency
	concurrencyChan := make(chan struct{}, maxConcurrency)

	for _, url := range urls {
		wg.Add(1)
		concurrencyChan <- struct{}{}
		go func(url string) {
			defer func() {
				<-concurrencyChan
			}()
			makeRequest(&wg, url)
		}(url)
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}

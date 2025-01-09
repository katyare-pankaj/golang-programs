package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchContent(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	duration := time.Since(start)
	fmt.Printf("Fetched %s in %s\n", url, duration)
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.wikipedia.org",
	}

	var wg sync.WaitGroup

	startTime := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go fetchContent(url, &wg)
	}

	wg.Wait()

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal processing time for all requests: %s\n", totalTime)
}

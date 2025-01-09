package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	duration := time.Since(start)
	fmt.Printf("Fetched %s in %s\n", url, duration)
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.wikipedia.org",
	}

	var wg sync.WaitGroup

	startTime := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg)
	}

	wg.Wait()

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal processing time: %s\n", totalTime)
}

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkStatus(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%s: %d\n", url, resp.StatusCode)
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.wikipedia.org",
		"https://www.non-existing-website.com",
	}

	var wg sync.WaitGroup

	startTime := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go checkStatus(url, &wg)
	}

	wg.Wait()

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal processing time: %s\n", totalTime)
}

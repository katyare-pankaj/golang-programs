package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkStatusCode(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %v\n", url, err)
		return
	}

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %v\n", url, err)
		return
	}

	statusCode := resp.StatusCode
	defer resp.Body.Close()

	duration := time.Since(start)
	result := fmt.Sprintf("URL: %s, Status Code: %d, Duration: %s\n", url, statusCode, duration)
	results <- result
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
	results := make(chan<- string)

	for _, url := range urls {
		wg.Add(1)
		go checkStatusCode(url, &wg, results)
	}

	go func() {
		for result := range results {
			fmt.Print(result)
		}
	}()

	wg.Wait()
	close(results)
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func scrapeURL(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done() // Signal completion of this task

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error scraping %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: %v\n", url, err)
		return
	}

	// Simulate processing the scraped data and add to results
	result := string(body)
	results <- result
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 100) // Buffered channel to handle multiple results

	urls := []string{
		"https://example.com",
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
	}

	// Start scraping tasks
	for _, url := range urls {
		wg.Add(1) // Increment the wait group for each task
		go scrapeURL(url, &wg, results)
	}

	// Wait for all scraping tasks to complete
	wg.Wait()

	// Close the results channel
	close(results)

	// Process all the results
	fmt.Println("Scraped data:")
	for result := range results {
		fmt.Println(result)
	}
}

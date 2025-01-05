package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func scrapeURL(url string, results chan<- string) {
	defer close(results) // Close the channel after sending the result

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

	// Simulate processing the scraped data and send the result
	result := string(body)
	results <- result
}

func consumeResults(wg *sync.WaitGroup, results <-chan string) {
	defer wg.Done() // Signal completion of this task

	fmt.Println("Consuming results:")
	for result := range results {
		fmt.Println(result)
	}
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://example.com",
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
	}

	// Launch a consumer goroutine
	wg.Add(1)
	go consumeResults(&wg, make(chan string, 100)) // Buffered channel to handle multiple results

	// Start scraping tasks
	for _, url := range urls {
		wg.Add(1)
		go scrapeURL(url, make(chan string, 100))
	}

	// Wait for all scraping tasks to complete
	wg.Wait()
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type ScrapeResult struct {
	URL     string
	Content string
}

func scrapeURL(url string, wg *sync.WaitGroup, results chan<- ScrapeResult) {
	defer wg.Done()

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

	results <- ScrapeResult{URL: url, Content: string(body)}
}

func main() {
	var wg sync.WaitGroup
	results := make(chan ScrapeResult, 100)

	urls := []string{
		"https://www.example.com",
		"https://www.example.com/page1",
		"https://www.example.com/page2",
		// Add more URLs as needed
	}

	// Start scraping tasks
	for _, url := range urls {
		wg.Add(1)
		go scrapeURL(url, &wg, results)
	}

	// Wait for all scraping tasks to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process the results
	fmt.Println("Scraped Data:")
	for result := range results {
		fmt.Printf("URL: %s\nContent: %s\n\n", result.URL, result.Content)
	}
}

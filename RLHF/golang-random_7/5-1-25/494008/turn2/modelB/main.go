package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const numWorkers = 10
const numUrls = 50

// worker function to scrape a single URL
func scrapeURL(url string, wg *sync.WaitGroup, results chan<- string) {
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
	results <- string(body)
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 100)
	// Generate sample URLs
	urls := make([]string, numUrls)
	for i := 0; i < numUrls; i++ {
		urls[i] = fmt.Sprintf("https://example.com/page%d", i)
	}

	// Start workers
	for i := 0; i < numWorkers; i++ {
		go func() {
			for url := range urls {
				scrapeURL(url, &wg, results)
			}
		}()
	}

	// Close the URL channel after all workers are started
	close(urls)

	// Wait for all workers to complete
	wg.Wait()
	close(results)

	// Process all the results
	fmt.Println("Scraped data:")
	for result := range results {
		fmt.Println(result)
	}
}

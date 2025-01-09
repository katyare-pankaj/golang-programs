package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Website represents a URL and its processing time.
type Website struct {
	URL      string
	Duration time.Duration
}

// FetchWebsite fetches data from a given website and updates the Duration field in Website.
func FetchWebsite(website *Website, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	_, err := http.Get(website.URL)
	if err != nil {
		website.Duration = -time.Since(start) // Negative duration indicates error
		fmt.Printf("Error fetching %s: %v\n", website.URL, err)
		return
	}
	website.Duration = time.Since(start)
	fmt.Printf("Fetched %s in %s\n", website.URL, website.Duration)
}

func main() {
	// Define websites to fetch data from
	websites := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.wikipedia.org",
		"https://example.com/non-existent-page", // This will fail to fetch
	}

	// Convert website URLs to Website struct to keep track of processing time
	websiteList := make([]*Website, len(websites))
	for i, url := range websites {
		websiteList[i] = &Website{URL: url}
	}

	var wg sync.WaitGroup
	startTime := time.Now()

	// Launch concurrent fetch operations for each website
	wg.Add(len(websiteList))
	for _, website := range websiteList {
		go FetchWebsite(website, &wg)
	}

	wg.Wait()

	// Calculate total processing time
	totalTime := time.Since(startTime)

	// Display results
	fmt.Printf("\nTotal processing time: %s\n", totalTime)
	fmt.Println("\nWebsite Processing Times:")
	for _, website := range websiteList {
		if website.Duration < 0 {
			fmt.Printf("Error fetching: %s\n", website.URL)
		} else {
			fmt.Printf("%s: %s\n", website.URL, website.Duration)
		}
	}
}

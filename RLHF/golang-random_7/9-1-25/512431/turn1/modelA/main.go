package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Resource represents a resource to fetch
type Resource struct {
	URL string
}

// fetchData fetches data from a given URL
func fetchData(wg *sync.WaitGroup, resource *Resource, data chan string) {
	defer wg.Done() // Signal that this goroutine has completed

	start := time.Now()
	res, err := http.Get(resource.URL)
	if err != nil {
		data <- fmt.Sprintf("Error fetching %s: %v\n", resource.URL, err)
		return
	}
	defer res.Body.Close()

	body, err := http.ReadAll(res.Body)
	if err != nil {
		data <- fmt.Sprintf("Error reading body from %s: %v\n", resource.URL, err)
		return
	}

	// For simplicity, we just print the body length here
	data <- fmt.Sprintf("Fetched %d bytes from %s in %s\n", len(body), resource.URL, time.Since(start))
}

func main() {
	resources := []Resource{
		{"https://www.example.com/resource1"},
		{"https://www.example.com/resource2"},
		{"https://www.example.com/resource3"},
		// Add more resources as needed
	}

	dataChannel := make(chan string)
	var wg sync.WaitGroup

	// Start a goroutine for each resource
	for _, resource := range resources {
		wg.Add(1) // Increment the wait group counter
		go fetchData(&wg, &resource, dataChannel)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the data channel
	close(dataChannel)

	// Output all fetched data
	for msg := range dataChannel {
		fmt.Println(msg)
	}

	// Calculate total processing time
	totalProcessingTime := time.Now()
	fmt.Printf("Total processing time: %s\n", totalProcessingTime)
}

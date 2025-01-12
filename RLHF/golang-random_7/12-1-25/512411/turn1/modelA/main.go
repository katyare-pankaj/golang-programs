package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Fetched %s\n", url)
	time.Sleep(time.Second) // Simulate processing time
}

func main() {
	var wg sync.WaitGroup

	urls := []string{"http://example.com/api/1", "http://example.com/api/2", "http://example.com/api/3"}

	// Start concurrent API calls
	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg)
	}

	// Wait for all API calls to complete
	wg.Wait()

	fmt.Println("All API calls have completed.")
}

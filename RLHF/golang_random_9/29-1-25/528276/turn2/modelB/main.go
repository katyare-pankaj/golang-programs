package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		// Deferred body close in the following defer statement.
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	// Process data...
	fmt.Printf("Fetched %s: %s\n", url, data)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://example.com", // Valid
		"https://invalid.com", // Invalid
	}
	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg)
	}
	wg.Wait()
	fmt.Println("All fetching operations complete.")
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Result struct {
	Data string `json:"data"`
}

func fetchData(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Make an HTTP GET request to the API URL
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data from %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to get data from %s: Status code %d\n", url, resp.StatusCode)
		return
	}

	// Decode the response body into a Result struct
	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding response from %s: %v\n", url, err)
		return
	}

	// Print the result
	fmt.Printf("Successfully fetched data from %s: %s\n", url, result.Data)
}

func main() {
	// Define some API URLs
	urls := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
		"https://api.example.com/data4",
		// Add more URLs as needed
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Start fetching data concurrently
	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All API calls completed.")
}

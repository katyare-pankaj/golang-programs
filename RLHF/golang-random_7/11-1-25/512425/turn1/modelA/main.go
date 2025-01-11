package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

// APIClient struct encapsulates the API client logic
type APIClient struct {
	baseURL string
}

// NewAPIClient initializes a new APIClient
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{baseURL: baseURL}
}

// makeAPIRequest performs an API call and returns the data and error
func (c *APIClient) makeAPIRequest(endpoint string) (data []byte, err error) {
	url := c.baseURL + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}
	return data, nil
}

// getDataFromAPI makes concurrent API requests
func (c *APIClient) getDataFromAPI(endpoints []string, numGoroutines int) (results [][]byte, err error) {
	var wg sync.WaitGroup
	results = make([][]byte, len(endpoints))
	for i, endpoint := range endpoints {
		wg.Add(1)
		go func(i int, endpoint string) {
			defer wg.Done()
			data, err := c.makeAPIRequest(endpoint)
			if err != nil {
				log.Printf("Error retrieving data from %s: %v\n", endpoint, err)
				results[i] = nil
			} else {
				results[i] = data
			}
		}(i, endpoint)
	}
	wg.Wait()
	return results, nil
}

func main() {
	baseURL := "https://api.example.com/v1/"
	endpoints := []string{"vehicles", "routes", "stations"}
	numGoroutines := len(endpoints)

	client := NewAPIClient(baseURL)
	results, err := client.getDataFromAPI(endpoints, numGoroutines)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for i, endpoint := range endpoints {
		if results[i] != nil {
			fmt.Printf("Data from %s:\n%s\n", endpoint, string(results[i]))
		} else {
			fmt.Printf("Failed to retrieve data from %s.\n", endpoint)
		}
	}
}

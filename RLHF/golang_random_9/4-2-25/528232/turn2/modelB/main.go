package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func fetchData(url string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close() // Ensure response body is closed

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func main() {
	url := "https://api.example.com/data"
	data, err := fetchData(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	fmt.Println("Response from API:", data)
}

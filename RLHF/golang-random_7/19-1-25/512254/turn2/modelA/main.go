package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

// Define a function to perform HTTP requests
func makeRequest(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Function to perform an HTTP request with retries and circuit breaker
func executeRequestWithRetry(url string, retries int, backoff time.Duration, cb *gobreaker.CircuitBreaker) (*http.Response, error) {
	var response *http.Response
	var err error

	for i := 0; i < retries; i++ {
		log.Printf("Attempt %d for %s\n", i+1, url)

		// Execute using the circuit breaker
		result, err := cb.Execute(func() (interface{}, error) {
			return makeRequest(url)
		})

		if err == nil {
			response = result.(*http.Response)
			return response, nil
		}

		log.Printf("Error: %v\n", err)
		time.Sleep(backoff)

		// Increase backoff interval exponentially
		backoff *= 2
	}

	return response, err
}

func main() {
	// Configure the circuit breaker settings
	cbSettings := gobreaker.Settings{
		Name:        "HTTP GET Circuit Breaker",
		MaxRequests: 5,
		Interval:    5 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("State changed in %s from %s to %s", name, from, to)
		},
	}
	cb := gobreaker.NewCircuitBreaker(cbSettings)

	url := "http://example.com"
	retries := 3
	initialBackoff := 2 * time.Second

	resp, err := executeRequestWithRetry(url, retries, initialBackoff, cb)
	if err != nil {
		log.Fatalf("Failed to get response: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v\n", err)
	}

	log.Printf("Successfully fetched: %s\n", resp.Status)
	fmt.Println("Response Body:", string(body))
}

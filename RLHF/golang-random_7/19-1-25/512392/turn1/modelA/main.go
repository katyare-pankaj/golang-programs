package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

// makeRequest represents a sample HTTP request operation
func makeRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// executeRequestWithRetry handles retries with a simple backoff strategy
func executeRequestWithRetry(url string, retries int, backoff time.Duration, cb *gobreaker.CircuitBreaker) (*http.Response, error) {
	var response *http.Response
	var err error

	for i := 0; i < retries; i++ {
		log.Printf("Attempt %d for %s\n", i+1, url)

		result, err := cb.Execute(func() (interface{}, error) {
			return makeRequest(url)
		})

		if err == nil {
			return result.(*http.Response), nil
		}

		log.Println("Error:", err)
		time.Sleep(backoff)

		// Exponential backoff
		backoff *= 2
	}

	return response, err
}

func main() {
	// Configure the circuit breaker
	cbSettings := gobreaker.Settings{
		Name:        "HTTP GET Circuit Breaker",
		MaxRequests: 5,
		Interval:    5 * time.Second,
		Timeout:     60 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("State changed from %s to %s", from, to)
		},
	}
	cb := gobreaker.NewCircuitBreaker(cbSettings)

	url := "http://example.com"
	retries := 3
	backoff := 1 * time.Second

	resp, err := executeRequestWithRetry(url, retries, backoff, cb)
	if err != nil {
		log.Fatalf("Failed to get response: %v\n", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Successfully fetched: %s\n", resp.Status)
}

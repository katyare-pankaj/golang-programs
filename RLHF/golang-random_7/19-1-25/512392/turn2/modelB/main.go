package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/sony/gobreaker"
)

const (
	httpMaxRetries   = 3
	httpRetryBackoff = 1 * time.Second
)

// CircuitBreakerClient encapsulates the HTTP client with circuit breaker
type CircuitBreakerClient struct {
	client *http.Client
	cb     *gobreaker.CircuitBreaker
}

// NewCircuitBreakerClient creates a new CircuitBreakerClient
func NewCircuitBreakerClient() *CircuitBreakerClient {
	cbSettings := gobreaker.Settings{
		Name:        "CircuitBreakerClient",
		MaxRequests: 10,
		Interval:    5 * time.Second,
		Timeout:     15 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Trip the circuit if there are more than 50% failed requests in the last 5 seconds
			return counts.ConsecutiveFailures > 5
		},
	}
	cb := gobreaker.NewCircuitBreaker(cbSettings)
	return &CircuitBreakerClient{
		client: &http.Client{Timeout: 10 * time.Second},
		cb:     cb,
	}
}

// Do executes a HTTP request with circuit breaker and retries
func (c *CircuitBreakerClient) Do(req *http.Request) (*http.Response, error) {
	// Wrap the HTTP request in a retryable function
	retryableFn := func() (interface{}, error) {
		res, err := c.client.Do(req)
		if err != nil {
			return nil, errors.Wrap(err, "failed to make HTTP request")
		}
		defer res.Body.Close()

		// Check for 5xx status codes
		if res.StatusCode >= 500 && res.StatusCode < 600 {
			return nil, fmt.Errorf("server error: %d", res.StatusCode)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read response body")
		}

		return &http.Response{
			StatusCode: res.StatusCode,
			Header:     res.Header,
			Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
		}, nil
	}

	// Execute the retryable function with the circuit breaker
	result, err := c.cb.Execute(func() (interface{}, error) {
		return retryableFn()
	})

	if err != nil {
		// Retries exhausted, return the last error
		return nil, err
	}

	return result.(*http.Response), nil
}

func main() {
	client := NewCircuitBreakerClient()

	u, err := url.Parse("http://example.com/api/data")
	if err != nil {
		log.Fatal("Failed to parse URL:", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, u.String(), nil)
	if err != nil {
		log.Fatal("Failed to create HTTP request:", err)
	}

	for i := 1; i <= 5; i++ {
		log.Printf("Attempt %d...", i)
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		defer resp.Body.Close()
		log.Println("Response status:", resp.StatusCode)
		if resp.StatusCode == http.StatusOK {
			break
		}
		time.Sleep(httpRetryBackoff)
	}
}

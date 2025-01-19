package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/sony/gobreaker"
)

const (
	defaultRetryAttempts  = 3
	defaultRetryTimeout   = 5 * time.Second
	circuitBreakerTimeout = 60 * time.Second
)

var (
	errServerUnavailable = fmt.Errorf("server unavailable")
	client               = http.Client{Timeout: 10 * time.Second}
)

func callServer(url string) error {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

func retryCallServer(url string) error {
	return retry.Do(func() error {
		return callServer(url)
	},
		retry.Attempts(defaultRetryAttempts),
		retry.DelayType(retry.FixedDelay),
		retry.Delay(defaultRetryTimeout),
	)
}

func circuitBreakerCallServer(url string) error {
	breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "server-breaker",
		Timeout: circuitBreakerTimeout,
	})

	return breaker.Run(func() error {
		return retryCallServer(url)
	})
}

func main() {
	serverURL := "http://example.com/some-endpoint"
	err := circuitBreakerCallServer(serverURL)
	if err != nil {
		log.Printf("Error calling server: %v\n", err)
		if gobreaker.IsOpen(err) {
			log.Println("Circuit breaker is open. Request will be blocked.")
		}
	} else {
		log.Println("Call to server successful.")
	}
}

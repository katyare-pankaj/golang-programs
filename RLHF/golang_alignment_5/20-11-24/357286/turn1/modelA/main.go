package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxRetries       = 3
	openThreshold    = 3
	closeThreshold   = 10
	halfOpenInterval = 10 * time.Second
)

type CircuitBreaker struct {
	state            string
	failures         int
	successes        int
	totalAttempts    int
	lastOpenTime     time.Time
	lastHalfOpenTime time.Time
	mtx              sync.RWMutex
}

func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{state: "Closed"}
}

func (cb *CircuitBreaker) attempt(ctx context.Context, f func(context.Context) error) error {
	cb.mtx.RLock()
	state := cb.state
	cb.mtx.RUnlock()

	switch state {
	case "Open":
		return fmt.Errorf("circuit breaker is open")
	case "Half-Open":
		if time.Now().Sub(cb.lastHalfOpenTime) < halfOpenInterval {
			return fmt.Errorf("circuit breaker is in half-open state")
		}
		cb.mtx.Lock()
		cb.state = "Closed"
		cb.mtx.Unlock()
		return cb.attempt(ctx, f)
	}

	err := f(ctx)
	if err != nil {
		cb.mtx.Lock()
		cb.failures++
		cb.totalAttempts++
		if cb.failures >= openThreshold {
			cb.state = "Open"
			cb.lastOpenTime = time.Now()
		}
		cb.mtx.Unlock()
		return err
	}

	cb.mtx.Lock()
	cb.successes++
	cb.totalAttempts++
	if cb.failures == 0 {
		cb.lastHalfOpenTime = time.Now()
	}
	cb.mtx.Unlock()
	return nil
}

func (cb *CircuitBreaker) retry(ctx context.Context, f func(context.Context) error, delay time.Duration) error {
	for attempt := 0; attempt < maxRetries; attempt++ {
		err := cb.attempt(ctx, f)
		if err == nil {
			return nil
		}
		if attempt < maxRetries-1 {
			fmt.Println("Attempt", attempt+1, "failed, retrying in", delay)
			time.Sleep(delay)
		}
	}
	return fmt.Errorf("all retries failed")
}

func externalService(ctx context.Context) error {
	// Simulate an error with a certain probability
	if rand.Intn(10) < 6 {
		return fmt.Errorf("external service failed")
	}
	return nil
}

func main() {
	ctx := context.Background()
	cb := NewCircuitBreaker()

	for i := 0; i < 20; i++ {
		err := cb.retry(ctx, externalService, 1*time.Second)
		if err != nil {
			fmt.Println("Final error:", err)
		}
	}
}

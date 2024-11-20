package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CircuitBreaker struct {
	mutex             sync.Mutex
	state             string
	failedAttempts    int
	maxFailedAttempts int
	resetTime         time.Duration
	lastFailedAttempt time.Time
	retryDelay        time.Duration
}

func NewCircuitBreaker(maxFailedAttempts int, resetTime time.Duration, retryDelay time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:             "closed",
		maxFailedAttempts: maxFailedAttempts,
		resetTime:         resetTime,
		retryDelay:        retryDelay,
	}
}

func (cb *CircuitBreaker) Allow() bool {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.state == "open" {
		if time.Since(cb.lastFailedAttempt) > cb.resetTime {
			cb.state = "half-open"
			cb.failedAttempts = 0
		}
	}

	return cb.state != "open"
}

func (cb *CircuitBreaker) Trip() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.state != "open" {
		cb.state = "open"
		cb.lastFailedAttempt = time.Now()
		cb.failedAttempts++
	}
}

func (cb *CircuitBreaker) Reset() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.state == "half-open" {
		cb.state = "closed"
		cb.failedAttempts = 0
	}
}

func (cb *CircuitBreaker) RetryDelay() time.Duration {
	return cb.retryDelay
}

func main() {
	ctx := context.Background()
	cb := NewCircuitBreaker(3, 5*time.Second, 1*time.Second)

	doWork := func(ctx context.Context) error {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		if rand.Intn(10) < 3 {
			return errors.New("service failed")
		}
		fmt.Println("Work completed successfully.")
		return nil
	}

	retryDoWork := func(ctx context.Context, f func(context.Context) error) error {
		for {
			err := f(ctx)
			if err == nil {
				return nil
			}

			fmt.Println("Work failed. Retrying...")

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(cb.RetryDelay()):
			}
		}
	}

	for i := 0; i < 10; i++ {
		if cb.Allow() {
			if err := retryDoWork(ctx, doWork); err != nil {
				cb.Trip()
				fmt.Println("Circuit breaker tripped.")
			} else {
				cb.Reset()
			}
		} else {
			fmt.Println("Circuit breaker is open. Skipping work.")
		}
	}
}

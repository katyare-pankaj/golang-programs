package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Payment represents a payment request.
type Payment struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

// Recommender generates real-time recommendations.
type Recommender interface {
	Recommend(ctx context.Context, payment Payment) ([]string, error)
}

// Throttler controls the rate of recommendation requests.
type Throttler interface {
	Acquire(context.Context) error
	Release(context.Context)
}

// processPayment processes a payment request and generates recommendations.
func processPayment(ctx context.Context, payment Payment, recommender Recommender, throttler Throttler) error {
	defer throttler.Release(ctx)

	if err := throttler.Acquire(ctx); err != nil {
		return err
	}

	recommendations, err := recommender.Recommend(ctx, payment)
	if err != nil {
		return err
	}

	fmt.Println("Payment", payment.ID, ": Recommendations:", recommendations)

	return nil
}

func main() {
	ctx := context.Background()

	// Replace with your actual implementation of Recommender
	recommender := &mockRecommender{}

	// Replace with your actual implementation of Throttler
	throttler := &mockThrottler{maxRate: 5, concurrency: 2}

	payments := []Payment{
		{ID: "1", Amount: 100.0},
		{ID: "2", Amount: 50.0},
		{ID: "3", Amount: 75.0},
		{ID: "4", Amount: 200.0},
		{ID: "5", Amount: 125.0},
	}

	var wg sync.WaitGroup

	for _, payment := range payments {
		wg.Add(1)
		go func(p Payment) {
			defer wg.Done()
			if err := processPayment(ctx, p, recommender, throttler); err != nil {
				fmt.Println("Error processing payment:", err)
			}
		}(payment)
	}

	wg.Wait()
}

// mockRecommender is a mock implementation of Recommender.
type mockRecommender struct {
}

func (m *mockRecommender) Recommend(ctx context.Context, payment Payment) ([]string, error) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return []string{"Product X", "Product Y", "Product Z"}, nil
}

// mockThrottler is a mock implementation of Throttler.
type mockThrottler struct {
	mu          sync.Mutex
	count       int
	maxRate     int
	concurrency int
}

func (m *mockThrottler) Acquire(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for m.count >= m.concurrency {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}

	m.count++
	return nil
}

func (m *mockThrottler) Release(ctx context.Context) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.count--
}

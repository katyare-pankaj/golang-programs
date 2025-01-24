package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// simulateDatabaseQuery simulates executing a database query, respecting context cancellation.
func simulateDatabaseQuery(ctx context.Context, id int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	// Simulate a query processing time
	queryDuration := time.Duration(2+id%3) * time.Second

	select {
	case <-time.After(queryDuration): // Simulate query taking some time
		ch <- fmt.Sprintf("Query %d completed successfully", id)
	case <-ctx.Done():
		ch <- fmt.Sprintf("Query %d cancelled due to timeout", id)
	}
}

func main() {
	const timeout = 3 * time.Second
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Channel to capture results from each query
	results := make(chan string, 5)

	// Simulate 5 concurrent database queries
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go simulateDatabaseQuery(ctx, i, &wg, results)
	}

	// Wait for all queries to complete or be cancelled
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print results
	for result := range results {
		fmt.Println(result)
	}
}

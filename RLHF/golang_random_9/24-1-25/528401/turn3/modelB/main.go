package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// simulateQuery simulates a database query that takes a random amount of time.
// It respects the context cancellation if the query exceeds the timeout.
func simulateQuery(ctx context.Context, id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	// Generate a random sleep duration between 0.5 and 2 seconds
	duration := time.Duration(rand.Intn(1500)+500) * time.Millisecond

	select {
	case <-time.After(duration):
		results <- fmt.Sprintf("Query %d completed successfully after %s", id, duration)
	case <-ctx.Done():
		results <- fmt.Sprintf("Query %d cancelled due to timeout", id)
	}
}

func main() {
	const numQueries = 5            // Number of queries to simulate
	const timeout = 2 * time.Second // Timeout for each query

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	results := make(chan string, numQueries) // Channel to collect query results

	// Start goroutines to simulate queries
	wg.Add(numQueries)
	for i := 1; i <= numQueries; i++ {
		go simulateQuery(ctx, i, &wg, results)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Receive and print the query results
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All queries processed.")
}

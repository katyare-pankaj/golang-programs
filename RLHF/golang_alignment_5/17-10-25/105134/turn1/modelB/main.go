package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to simulate an external API call that takes time to complete.
func externalAPIcall(ctx context.Context, result chan<- string, index int) {
	select {
	case <-ctx.Done():
		fmt.Printf("Request %d cancelled\n", index)
		return
	default:
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		result <- fmt.Sprintf("Result from request %d", index)
	}
}

// Function to fan-out multiple requests to external APIs and fan-in the results.
func fanOutFanIn(ctx context.Context, numRequests int) []string {
	results := make(chan string, numRequests)
	var wg sync.WaitGroup

	// Fan-out: Start multiple goroutines for each request
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			externalAPIcall(ctx, results, i)
		}(i)
	}

	// Fan-in: Collect results from the channels into a slice
	var finalResults []string
	for i := 0; i < numRequests; i++ {
		select {
		case result := <-results:
			finalResults = append(finalResults, result)
		case <-ctx.Done():
			fmt.Println("Context cancelled, returning partial results")
			return finalResults
		}
	}

	return finalResults
}

func main() {
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	numRequests := 10
	results := fanOutFanIn(ctx, numRequests)

	fmt.Println("Final Results:")
	for _, result := range results {
		fmt.Println(result)
	}
}

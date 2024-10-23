package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate a database query function
func queryDatabase(query string) string {
	time.Sleep(time.Millisecond * 500) // Simulate query time
	return fmt.Sprintf("Result for query: %s", query)
}

// Handle a single query request concurrently
func handleQueryRequest(query string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	result := queryDatabase(query)
	results <- result
}

func main() {
	queries := []string{
		"SELECT * FROM users",
		"SELECT * FROM posts",
		"SELECT * FROM comments",
		// Add more queries here
	}

	start := time.Now()

	// Concurrent processing using goroutines and channels
	wg := &sync.WaitGroup{}
	results := make(chan string, len(queries))

	for _, query := range queries {
		wg.Add(1)
		go handleQueryRequest(query, wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print results
	for result := range results {
		fmt.Println(result)
	}

	fmt.Printf("Execution time: %s\n", time.Since(start))
}

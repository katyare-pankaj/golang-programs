package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Handler to simulate processing some long-running task.
func worker(ctx context.Context, wg *sync.WaitGroup, id int, ch chan<- string) {
	defer wg.Done()

	// Simulating a heavy task with a variable duration
	select {
	case <-time.After(time.Duration((2+id)%3) * time.Second): // Random long task
		ch <- fmt.Sprintf("Task %d completed", id)
	case <-ctx.Done():
		ch <- fmt.Sprintf("Task %d cancelled", id)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	const timeout = 2 * time.Second
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	results := make(chan string, 3)
	wg.Add(3) // Assume we are kicking off 3 worker tasks

	// Start three tasks
	go worker(ctx, &wg, 1, results)
	go worker(ctx, &wg, 2, results)
	go worker(ctx, &wg, 3, results)

	// Wait for tasks or timeout
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Fprintln(w, result)
	}
}

func main() {
	http.HandleFunc("/process", apiHandler)
	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

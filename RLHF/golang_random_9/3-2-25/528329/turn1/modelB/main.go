package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Define a struct to hold the API's configuration
type config struct {
	Port           string
	RequestTimeout time.Duration
}

// Create a new instance of the config struct
var cfg = config{
	Port:           ":8080",
	RequestTimeout: 5 * time.Second,
}

// Handler function that processes requests
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Create a wait group to manage concurrent tasks
	var wg sync.WaitGroup

	// Define the number of concurrent tasks
	numTasks := 3

	// Add tasks to the wait group
	wg.Add(numTasks)

	// Start each task in a separate goroutine
	for i := 0; i < numTasks; i++ {
		go func(i int) {
			defer wg.Done()
			// Simulate work with a context
			doWork(ctx, i)
		}(i)
	}

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Fprintf(w, "Hello, world!")
}

// Simulate work being done by each goroutine
func doWork(ctx context.Context, taskNum int) {
	// Add a delay to simulate work
	time.Sleep(2 * time.Second)

	// Check if the context has timed out
	select {
	case <-ctx.Done():
		// Context has timed out, handle it
		fmt.Printf("Task %d timed out\n", taskNum)
		return
	default:
		// Context is not timed out, continue processing
		fmt.Printf("Task %d completed\n", taskNum)
	}
}

func main() {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), cfg.RequestTimeout)
	defer cancel()

	// Serve the HTTP server with the timeout context
	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: http.HandlerFunc(handler),
		BaseContext: func(_ *http.Request) context.Context {
			return ctx
		},
	}

	// Start the HTTP server
	fmt.Printf("Server listening on port %s\n", cfg.Port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

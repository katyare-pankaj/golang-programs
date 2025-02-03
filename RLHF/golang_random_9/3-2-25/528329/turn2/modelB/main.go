package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	http.Server
}

type response struct {
	Message string `json:"message"`
}

// SimulateWork represents the work that each goroutine will perform.
// It should accept a context.Context to support cancellation.
func SimulateWork(ctx context.Context, workID int) error {
	r := rand.Intn(10) + 1 // Random time between 1 to 10 seconds
	dur := time.Duration(r) * time.Second

	fmt.Printf("Worker %d: Starting work with duration %v...\n", workID, dur)
	select {
	case <-time.After(dur):
		// Simulate work completion
		fmt.Printf("Worker %d: Work completed.\n", workID)
		return nil
	case <-ctx.Done():
		// Handle context cancellation
		fmt.Printf("Worker %d: Context canceled, aborting work.\n", workID)
		return ctx.Err()
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Create a context with a global timeout of 10 seconds
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	const numWorkers = 5 // Number of concurrent tasks
	var wg sync.WaitGroup
	errs := make([]error, numWorkers)

	// Start the workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			// Create a context with a per-worker timeout of 5 seconds (e.g., to prevent individual worker from hogging the resource)
			workerCtx, workerCancel := context.WithTimeout(ctx, 5*time.Second)
			defer workerCancel()

			// Execute the work with the per-worker context
			errs[workerID] = SimulateWork(workerCtx, workerID)
		}(i)
	}

	// Wait for all workers to complete or context timeout to occur
	go func() {
		wg.Wait()
		close(make(chan struct{}))
	}()

	select {
	case <-make(chan struct{}):
		// All workers have finished successfully or have been canceled by timeout.
	}

	var results []response
	for i, err := range errs {
		if err != nil {
			// Handle worker error (e.g., context canceled, time out)
			if err == context.DeadlineExceeded {
				results = append(results, response{Message: fmt.Sprintf("Worker %d timed out", i)})
			} else {
				results = append(results, response{Message: fmt.Sprintf("Worker %d failed: %v", i, err)})
			}
		} else {
			results = append(results, response{Message: fmt.Sprintf("Worker %d completed successfully", i)})
		}
	}

	// Send the response to the client
	sendResponse(w, results, http.StatusOK)
}

func sendResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	srv := &Server{
		Server: http.Server{
			Addr: ":8080",
		},
	}
	http.HandleFunc("/api", apiHandler)
	log.Println("Starting server on :8080")
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type response struct {
	Results []string `json:"results"`
}

func main() {
	http.HandleFunc("/api", apiHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	results := make(chan string, 3) // Buffer size should match the number of tasks
	errors := make(chan error, 3)

	tasks := []func(context.Context) (string, error){
		task1,
		task2,
		task3,
	}

	// Launch each task in its own goroutine with its own context
	for _, task := range tasks {
		wg.Add(1)
		go func(t func(context.Context) (string, error)) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
			defer cancel()

			select {
			case <-ctx.Done():
				// Context was canceled or timed out
				if ctx.Err() == context.DeadlineExceeded {
					errors <- ctx.Err()
				}
			default:
				if result, err := t(ctx); err != nil {
					errors <- err
				} else {
					results <- result
				}
			}
		}(task)
	}

	// Close channels once all goroutines complete
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	var finalResults []string
	for result := range results {
		finalResults = append(finalResults, result)
	}

	// Send the response
	if len(errors) > 0 {
		http.Error(w, "Some tasks did not complete", http.StatusPartialContent)
		log.Println("Some tasks did not complete")
	} else {
		sendResponse(w, &response{Results: finalResults}, http.StatusOK)
		log.Println("All tasks completed successfully")
	}
}

func task1(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Second):
		return "Task 1 completed", nil
	}
}

func task2(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(4 * time.Second): // This task exceeds its timeout
		return "Task 2 completed", nil
	}
}

func task3(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(2 * time.Second):
		return "Task 3 completed", nil
	}
}

func sendResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

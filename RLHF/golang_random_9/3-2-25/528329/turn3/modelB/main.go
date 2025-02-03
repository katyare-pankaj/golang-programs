package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

type taskResult struct {
	Result string
	Error  error
}

func main() {
	srv := &Server{
		Server: http.Server{
			Addr: ":8080",
		},
	}
	http.HandleFunc("/api", apiHandler)
	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Define the tasks with their respective timeout durations
	tasks := []struct {
		Task func(context.Context) (string, error)
		Timeout time.Duration
	}{
		{task1, 2 * time.Second},
		{task2, 3 * time.Second},
		{task3, 1 * time.Second},
	}

	resultChan := make(chan taskResult)

	// Start the tasks in goroutines with their own context
	for _, task := range tasks {
		ctx, cancel := context.WithTimeout(r.Context(), task.Timeout)
		go func(ctx context.Context, task func(context.Context) (string, error)) {
			defer cancel()
			result, err := task(ctx)
			resultChan <- taskResult{Result: result, Error: err}
		}(ctx, task.Task)
	}

	var results []string
	var errors []error

	// Collect results and errors from the goroutines
	for range tasks {
		result := <-resultChan
		if result.Error != nil {
			errors = append(errors, result.Error)
		} else {
			results = append(results, result.Result)
		}
	}

	if len(errors) > 0 {
		// Handle any errors that occurred during task execution
		handleErrors(w, errors)
		return
	}

	sendResponse(w, &response{Message: "Tasks completed successfully: " + formatResults(results)}, http.StatusOK)
}

func handleErrors(w http.ResponseWriter, errors []error) {
	http.Error(w, "Some tasks failed: "+formatErrors(errors), http.StatusInternalServerError)
	log.Println("Tasks errors:", errors)
}

func formatErrors(errors []error) string {
	formatted := ""
	for _, err := range errors {
		formatted += err.Error() + "; "
	}
	return formatted
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
	case <-time.After(2 * time.Second):
		return "Task 2 completed", nil
	}
}

func task3(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
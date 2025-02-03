package main

import (
	"context"
	"encoding/json"
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
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	results := make(chan string, 3) // Buffer size should be equal to the number of tasks
	errChannel := make(chan error, 1)

	tasks := []func(context.Context) (string, error){
		task1,
		task2,
		task3,
	}

	for _, task := range tasks {
		wg.Add(1)
		go func(t func(context.Context) (string, error)) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				errChannel <- ctx.Err()
			default:
				if result, err := t(ctx); err != nil {
					errChannel <- err
				} else {
					results <- result
				}
			}
		}(task)
	}

	// Wait for all tasks to complete
	go func() {
		wg.Wait()
		close(results)
		close(errChannel)
	}()

	select {
	case <-ctx.Done():
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		log.Println("Request timed out")
		return
	case err := <-errChannel:
		if err != nil {
			http.Error(w, "Task error: "+err.Error(), http.StatusInternalServerError)
			log.Println("Task error:", err)
			return
		}
	}

	var finalResults []string
	for result := range results {
		finalResults = append(finalResults, result)
	}
	sendResponse(w, &response{Message: "Tasks completed successfully: " + formatResults(finalResults)}, http.StatusOK)
}

func task1(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(2 * time.Second):
		return "Task 1 completed", nil
	}
}

func task2(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(3 * time.Second):
		return "Task 2 completed", nil
	}
}

func task3(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Second):
		return "Task 3 completed", nil
	}
}

func sendResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func formatResults(results []string) string {
	formatted := ""
	for _, result := range results {
		formatted += result + "; "
	}
	return formatted
}

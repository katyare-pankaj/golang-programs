package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func longRunningTask(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	defer func() {
		log.Println("Long running task completed")
	}()

	select {
	case <-ctx.Done():
		log.Println("Task canceled by context")
		return
	default:
		time.Sleep(3 * time.Second)
	}

	fmt.Fprintf(w, "Task completed in 3 seconds")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go longRunningTask(ctx, w, r)

	select {
	case <-ctx.Done():
		// Context canceled, handle the error
		w.WriteHeader(http.StatusGatewayTimeout)
		fmt.Fprintf(w, "Request timed out")
		return
	case <-time.After(2 * time.Second):
		// If the task takes longer than the timeout, we can handle it here
		log.Println("Request completed in 2 seconds")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

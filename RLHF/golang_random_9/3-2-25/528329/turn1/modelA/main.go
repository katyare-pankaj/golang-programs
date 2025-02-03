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
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	done := make(chan struct{})

	// Increment WaitGroup counter
	wg.Add(1)

	go func() {
		defer wg.Done()
		select {
		case <-time.After(3 * time.Second):
			// Simulate work
			sendResponse(w, &response{Message: "Hello, World!"}, http.StatusOK)
		case <-ctx.Done():
			// Timeout or cancellation response
			return
		}
		close(done)
	}()

	// Wait for goroutine to complete or context timeout to occur
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return // Successful completion
	case <-ctx.Done():
		// Handle context canceled or deadline exceeded
		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("Request timed out")
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
			return
		}
		if ctx.Err() == context.Canceled {
			log.Printf("Request was canceled")
			http.Error(w, "Request was canceled", http.StatusBadRequest)
			return
		}
	}
}

func sendResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

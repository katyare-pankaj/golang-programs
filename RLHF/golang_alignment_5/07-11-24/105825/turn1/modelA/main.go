package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// IdempotentService represents an idempotent service
type IdempotentService struct {
	mu        sync.Mutex
	processed map[string]bool
}

// NewIdempotentService creates a new IdempotentService
func NewIdempotentService() *IdempotentService {
	return &IdempotentService{
		processed: make(map[string]bool),
	}
}

// ProcessRequest is an idempotent service endpoint
func (s *IdempotentService) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the request ID from the request body or headers, or generate a unique ID here
	reqID := r.FormValue("request_id")

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.processed[reqID] {
		fmt.Fprintf(w, "Request %s has already been processed\n", reqID)
		return
	}

	// Perform the idempotent operation here
	err := s.performIdempotentOperation(reqID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.processed[reqID] = true
	fmt.Fprintf(w, "Request %s has been processed successfully\n", reqID)
}

func (s *IdempotentService) performIdempotentOperation(reqID string) error {
	// Simulate an idempotent operation such as data processing, database updates, or sending messages
	// This operation should be safe to be called multiple times without causing any issues
	log.Printf("Processing request: %s", reqID)
	return nil
}

func main() {
	service := NewIdempotentService()
	http.HandleFunc("/process", service.ProcessRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// RequestData represents the data structure of a request
type RequestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ResponseData represents the data structure of a response
type ResponseData struct {
	Message string `json:"message"`
}

// IdempotencyStore is a simple in-memory store for idempotency keys
type IdempotencyStore struct {
	mu      sync.Mutex
	store   map[string]bool
	timeout time.Duration
}

// NewIdempotencyStore creates a new IdempotencyStore
func NewIdempotencyStore(timeout time.Duration) *IdempotencyStore {
	return &IdempotencyStore{
		store:   make(map[string]bool),
		timeout: timeout,
	}
}

// IsIdempotent checks if the request is idempotent and adds the key to the store
func (s *IdempotencyStore) IsIdempotent(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the key exists in the store
	if _, ok := s.store[key]; ok {
		return false // Request is not idempotent
	}

	// Add the key to the store
	s.store[key] = true

	// Go routine to remove the key after the timeout
	go func() {
		time.Sleep(s.timeout)
		s.mu.Lock()
		delete(s.store, key)
		s.mu.Unlock()
	}()

	return true
}

// idempotentHandler handles idempotent requests
func idempotentHandler(w http.ResponseWriter, r *http.Request, idempotencyStore *IdempotencyStore) {
	// Get the 'Idempotency-Key' header from the request
	idempotencyKey := r.Header.Get("Idempotency-Key")
	if idempotencyKey == "" {
		http.Error(w, "Idempotency-Key header is required", http.StatusBadRequest)
		return
	}

	// Check if the request is idempotent
	if !idempotencyStore.IsIdempotent(idempotencyKey) {
		http.Error(w, "Duplicate request detected", http.StatusConflict)
		return
	}

	// Process the request here
	var requestData RequestData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := ResponseData{
		Message: fmt.Sprintf("Request processed successfully: %s", requestData.Name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func main() {
	// Create an idempotency store with a timeout of 5 minutes
	idempotencyStore := NewIdempotencyStore(5 * time.Minute)

	http.HandleFunc("/idempotent", func(w http.ResponseWriter, r *http.Request) {
		idempotentHandler(w, r, idempotencyStore)
	})

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// For demonstration we'd use a mutex for state
var mu sync.Mutex

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	ProcessedMessage string `json:"processed_message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Defer close of request body to ensure cleanup
	defer r.Body.Close()

	// Mutex to ensure state consistency if necessary
	mu.Lock()
	defer mu.Unlock()

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Open a file to write the results
	file, err := os.OpenFile("output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	// Defer file closing
	defer file.Close()

	processedMessage := fmt.Sprintf("Processed: %s", req.Message)
	response := Response{ProcessedMessage: processedMessage}

	// Write response to file
	if _, err := file.WriteString(processedMessage + "\n"); err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/handler", handler)

	port := "8080"
	fmt.Printf("Server running locally on http://localhost:%s/handler\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

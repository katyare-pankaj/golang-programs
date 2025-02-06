package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// A simple struct to mimic a request payload
type Request struct {
	Message string `json:"message"`
}

// A simple struct to mimic a response
type Response struct {
	ProcessedMessage string `json:"processed_message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request body gets closed
	defer r.Body.Close()

	// Decode the incoming JSON payload
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Setup a log file for demo purposes
	logFile, err := os.OpenFile("serverless.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(w, "Could not open log file", http.StatusInternalServerError)
		return
	}
	// Ensure the log file gets closed
	defer logFile.Close()
	log.SetOutput(logFile)

	// Log the incoming message
	log.Printf("Received message: %s", req.Message)

	// Simulate processing
	processedMessage := fmt.Sprintf("Processed: %s", req.Message)

	// Respond with the processed message
	response := Response{ProcessedMessage: processedMessage}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Log the response
	log.Printf("Sent response: %s", processedMessage)
}

func main() {
	http.HandleFunc("/handler", handler)

	port := "8080"
	fmt.Printf("Serverless-like app running on http://localhost:%s/handler\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

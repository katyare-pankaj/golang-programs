package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Define a basic response structure
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Example handler function
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Example dynamic data
	var userData map[string]interface{}
	switch r.Method {
	case http.MethodGet:
		// Simulate fetching data from a database or API
		userData = map[string]interface{}{
			"id":    1,
			"name":  "John Doe",
			"email": "john@example.com",
		}
	case http.MethodPost:
		// Simulate processing a POST request
		userData = map[string]interface{}{
			"id":    2,
			"name":  "Jane Doe",
			"email": "jane@example.com",
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create a Response instance
	response := Response{
		Status:  "success",
		Message: "Request processed successfully.",
		Data:    userData,
	}

	// Encode the struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error writing JSON response: %v", err)
		return
	}
}

func main() {
	http.HandleFunc("/api/example", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

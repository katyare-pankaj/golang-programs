package main

import (
	"encoding/json"
	"net/http"
)

// Define the struct for the JSON response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Example data
	userData := map[string]interface{}{
		"id":    1,
		"name":  "John Doe",
		"email": "john@example.com",
	}

	// Create a Response instance
	response := Response{
		Status:  "success",
		Message: "User data retrieved successfully.",
		Data:    userData,
	}

	// Encode the struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the JSON response
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Define the HTTP route and handler
	http.HandleFunc("/getUserData", handleRequest)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}

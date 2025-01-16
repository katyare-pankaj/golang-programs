package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a struct to represent the JSON response
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Sample data for demonstration purposes
var users = map[int]string{
	1: "Alice",
	2: "Bob",
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request URL
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		// Handle invalid user ID
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Find the user in the sample data
	userName, ok := users[userID]
	if !ok {
		// Handle user not found
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	// Create the response data
	responseData := map[string]string{"user_name": userName}

	// Respond with success
	respondJSON(w, http.StatusOK, "User found successfully", responseData)
}

// Helper function to respond with JSON data
func respondJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	response := APIResponse{
		Status:  http.StatusText(status),
		Message: message,
		Data:    data,
	}

	// Set the content-type header
	w.Header().Set("Content-Type", "application/json")

	// Set the status code
	w.WriteHeader(status)

	// Encode the response as JSON
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		// Handle encoding error
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// Helper function to respond with an error
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, message, nil)
}

func main() {
	http.HandleFunc("/user", getUserHandler)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

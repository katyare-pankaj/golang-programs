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

// Helper function to validate and parse input parameters
func parseInputParameters(w http.ResponseWriter, r *http.Request, requiredParams []string) (map[string]string, error) {
	params := make(map[string]string)

	for _, param := range requiredParams {
		value := r.URL.Query().Get(param)
		if value == "" {
			return nil, fmt.Errorf("Missing required parameter: %s", param)
		}
		params[param] = value
	}

	return params, nil
}

// Handler function to get user details
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Define required parameters
	requiredParams := []string{"id"}

	// Parse input parameters
	params, err := parseInputParameters(w, r, requiredParams)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Parse the user ID
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, fmt.Sprintf("Invalid user ID: %s", params["id"]))
		return
	}

	// Find the user in the sample data
	userName, ok := users[userID]
	if !ok {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	// Create the response data
	responseData := map[string]string{"user_name": userName}

	// Respond with success
	respondJSON(w, http.StatusOK, "User found successfully", responseData)
}

func main() {
	http.HandleFunc("/user", getUserHandler)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

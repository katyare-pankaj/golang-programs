package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// APIResponse is a generic struct to hold API responses
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// User represents a user in the API
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// errorResponse is a helper struct to hold error responses
type errorResponse struct {
	Error string `json:"error"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request URL
	userID, err := parseUserID(r)
	if err != nil {
		respondError(w, err)
		return
	}

	// Find the user in the sample data
	user, err := findUser(userID)
	if err != nil {
		respondError(w, err)
		return
	}

	// Respond with success
	respondJSON(w, http.StatusOK, "User found successfully", user)
}

func parseUserID(r *http.Request) (int, error) {
	userIDStr := r.URL.Query().Get("id")
	if userIDStr == "" {
		return 0, fmt.Errorf("user ID is required")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID: %v", err)
	}

	return userID, nil
}

func findUser(userID int) (*User, error) {
	// Sample data lookup
	users := map[int]string{
		1: "Alice",
		2: "Bob",
	}

	userName, ok := users[userID]
	if !ok {
		return nil, fmt.Errorf("user not found with ID %d", userID)
	}

	return &User{ID: userID, Name: userName}, nil
}

func respondJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	response := APIResponse{
		Status:  http.StatusText(status),
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func respondError(w http.ResponseWriter, err error) {
	var errorMsg string
	if e, ok := err.(*json.UnmarshalTypeError); ok {
		// Handle specific JSON unmarshal errors for better error messages
		errorMsg = fmt.Sprintf("invalid value for %s: %s", e.Field, e.Value)
	} else {
		errorMsg = err.Error()
	}

	errResponse := errorResponse{Error: errorMsg}
	respondJSON(w, http.StatusBadRequest, "Error", errResponse)
}

func main() {
	http.HandleFunc("/user", getUserHandler)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

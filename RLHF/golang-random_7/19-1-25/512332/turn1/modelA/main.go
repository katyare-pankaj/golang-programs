package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

// Input represents the JSON structure expected in the API request.
type Input struct {
	Name string `json:"name"`
}

// validateInput performs basic validation on user input.
func validateInput(input Input) bool {
	// Use a regex to allow only letters and spaces (adjust pattern according to your needs)
	validName := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return validName.MatchString(input.Name)
}

// sanitizeOutput performs basic output escaping to prevent injection attacks.
func sanitizeOutput(input string) string {
	return template.HTMLEscapeString(input)
}

// callbackHandler is the HTTP handler for the API callback.
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input Input
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if !validateInput(input) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Sanitize the output before using it in responses
	sanitizedOutput := sanitizeOutput(input.Name)

	response := map[string]string{
		"message": "Hello, " + sanitizedOutput,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/callback", callbackHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

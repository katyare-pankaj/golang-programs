package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Input represents the JSON structure expected in the API request.
type Input struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

// validateAndSanitize performs validation and sanitization on user input.
func validateAndSanitize(input *Input) string {
	var messages []string

	// Validate Email with a regular expression
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(input.Email) {
		messages = append(messages, "Invalid email format.")
	}

	// Validate Phone Number allowing only digits and some common phone number characters
	phoneRegex := regexp.MustCompile(`^[\d\s()+-]+$`)
	if !phoneRegex.MatchString(input.PhoneNumber) {
		messages = append(messages, "Invalid phone number format.")
	}

	// Sanitize input to avoid injection attacks
	input.Email = template.HTMLEscapeString(input.Email)
	input.PhoneNumber = template.HTMLEscapeString(input.PhoneNumber)

	// Generate error messages if any validation failed
	if len(messages) > 0 {
		return strings.Join(messages, " ")
	}

	return ""
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

	// Perform validation and sanitization
	errorMessage := validateAndSanitize(&input)
	if errorMessage != "" {
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message": "Input successfully validated and sanitized",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/validate", callbackHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

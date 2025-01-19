package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// UserInput represents the JSON structure expected in the API request.
type UserInput struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// sanitizePhone sanitizes the phone number by removing any non-digit characters.
func sanitizePhone(phone string) string {
	return strings.Join(regexp.MustCompile(`\d`).FindAllString(phone, -1), "")
}

// validateName checks if the name contains only alphabets and spaces using a regular expression.
func validateName(name string) bool {
	const nameRegex = `^[a-zA-Z\s]+$`
	return regexp.MustCompile(nameRegex).MatchString(name)
}

// callbackHandler is the HTTP handler for the API callback.
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input UserInput
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Bad request: Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Trim any leading or trailing spaces from the name and phone number
	input.Name = strings.TrimSpace(input.Name)
	input.Phone = strings.TrimSpace(input.Phone)

	// Sanitize phone number
	sanitizedPhone := sanitizePhone(input.Phone)

	// Validate name
	if !validateName(input.Name) {
		http.Error(w, "Invalid name: Name can only contain alphabets and spaces", http.StatusBadRequest)
		return
	}
	response := map[string]string{
		"name":  input.Name,
		"phone": sanitizedPhone,
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

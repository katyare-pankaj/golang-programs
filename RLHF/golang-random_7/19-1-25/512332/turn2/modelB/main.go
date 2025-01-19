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
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// validateEmail checks if the email address is valid using a regular expression.
func validateEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

// sanitizePhone sanitizes the phone number by removing any non-digit characters.
func sanitizePhone(phone string) string {
	return strings.Join(regexp.MustCompile(`\d`).FindAllString(phone, -1), "")
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

	// Trim any leading or trailing spaces from the email and phone number
	input.Email = strings.TrimSpace(input.Email)
	input.Phone = strings.TrimSpace(input.Phone)

	// Validate email
	if !validateEmail(input.Email) {
		http.Error(w, "Invalid email address", http.StatusBadRequest)
		return
	}

	// Sanitize phone number
	sanitizedPhone := sanitizePhone(input.Phone)
	if len(sanitizedPhone) < 10 || len(sanitizedPhone) > 15 {
		http.Error(w, "Invalid phone number length", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"email": input.Email,
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

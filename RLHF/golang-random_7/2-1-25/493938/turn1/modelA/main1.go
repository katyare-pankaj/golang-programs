package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// UserContext represents user-specific data
type UserContext struct {
	UserID   int
	Language string
}

// CustomError contains a message and additional fields for customization
type CustomError struct {
	Message  string `json:"message"`
	UserID   int    `json:"user_id"`
	Language string `json:"language"`
}

// FormatError formats a custom error message using fmt.Sprintf
func FormatError(template string, args ...interface{}) *CustomError {
	return &CustomError{
		Message: fmt.Sprintf(template, args...),
	}
}

// HandleErrorJSON responds with a JSON encoded CustomError
func HandleErrorJSON(w http.ResponseWriter, r *http.Request, err *CustomError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&CustomError{Message: "An unknown error occurred"})
	}
}

// GetUserContext retrieves user context from the request
func GetUserContext(r *http.Request) *UserContext {
	// For this example, assuming UserID and Language are derived from the request
	// Here we simulate extracting from request (headers or cookies)
	return &UserContext{
		UserID:   1,    // Example: fetch from session/cookies
		Language: "en", // Example: fetch from language header
	}
}

// MultilingualErrorMessage generates a multilingual error message
func MultilingualErrorMessage(template string, args ...interface{}) (string, error) {
	var messages = map[string]map[string]string{
		"en": {
			"generic":    "An error occurred.",
			"invalid_id": "The provided ID is invalid.",
		},
		"es": {
			"generic":    "Ocurrió un error.",
			"invalid_id": "El ID proporcionado es inválido.",
		},
	}

	// Get the user context from the request
	userContext := GetUserContext(nil) // Pass the `r` here, not `nil`
	lang := userContext.Language

	if lang == "" {
		lang = "en" // Default language
	}

	msg, ok := messages[lang][template]
	if !ok {
		return "", fmt.Errorf("unknown error message template: %s", template)
	}

	return fmt.Sprintf(msg, args...), nil
}

// ExampleHandler demonstrates how to handle a request with custom error messages
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")

	if !isValidUserID(userID) {
		// Conditional customization based on user ID
		customError, err := MultilingualErrorMessage("invalid_id", userID)
		if err != nil {
			log.Println(err)
			HandleErrorJSON(w, r, &CustomError{Message: "An unknown error occurred"})
			return
		}
		HandleErrorJSON(w, r, FormatError("Custom Error: %s", customError))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User ID %s is valid.", userID)
}

// isValidUserID simulates a validation function
func isValidUserID(userID string) bool {
	// Replace with actual validation logic
	return userID != ""
}

func main() {
	http.HandleFunc("/example", ExampleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

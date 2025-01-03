package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ErrorCode is an enum for different types of errors
type ErrorCode int

const (
	InvalidIDError    ErrorCode = 1
	UnauthorizedError           = 2
)

// ErrorResponse is a structured representation of an error
type ErrorResponse struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Context interface{} `json:"context,omitempty"`
}

// ErrorMessages contains error message templates by language
var ErrorMessages = map[string]map[ErrorCode]string{
	"en": {
		InvalidIDError:    "Invalid ID: %s",
		UnauthorizedError: "You are not authorized to perform this action.",
	},
	"es": {
		InvalidIDError:    "ID inválido: %s",
		UnauthorizedError: "No tiene permiso para realizar esta acción.",
	},
}

// GetErrorMessage retrieves an error message template by code and language
func GetErrorMessage(code ErrorCode, language string) string {
	msg, ok := ErrorMessages[language][code]
	if !ok {
		return ErrorMessages["en"][code] // Default to English if language not found
	}
	return msg
}

// FormatError formats an error message with context
func FormatError(code ErrorCode, language string, context interface{}) *ErrorResponse {
	msgTemplate := GetErrorMessage(code, language)
	formattedMsg := fmt.Sprintf(msgTemplate, context)
	return &ErrorResponse{
		Code:    code,
		Message: formattedMsg,
		Context: context,
	}
}

// HandleErrorJSON responds with a JSON encoded ErrorResponse
func HandleErrorJSON(w http.ResponseWriter, r *http.Request, err *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&ErrorResponse{Message: "An unknown error occurred"})
	}
}

// ExampleHandler demonstrates handling errors with conditional customization
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	// Example: Get user context from session or headers
	userContext := &UserContext{
		UserID:   1,
		Language: r.Header.Get("Accept-Language"), // Example: retrieve from header
		Role:     "admin",
	}

	userID := r.URL.Query().Get("userID")

	if !isValidUserID(userID) {
		err := FormatError(InvalidIDError, userContext.Language, userID)
		HandleErrorJSON(w, r, err)
		return
	}

	// Conditional customization: check user role
	if userContext.Role != "admin" {
		err := FormatError(UnauthorizedError, userContext.Language, nil)
		HandleErrorJSON(w, r, err)
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

// UserContext represents user-specific data
type UserContext struct {
	UserID   int
	Language string
	Role     string
}

func main() {
	http.HandleFunc("/example", ExampleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

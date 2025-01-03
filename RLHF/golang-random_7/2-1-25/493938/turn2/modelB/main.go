package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UserContext holds user context data, including language and user ID.
type UserContext struct {
	Language string
	UserID   int
}

// ErrorCustomiser is an interface that allows different types to implement
// custom error message handling.
type ErrorCustomiser interface {
	CustomiseError(template string, args ...interface{}) (string, error)
}

// customError represents a custom error with a message and contextual details.
type customError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// NewCustomError creates a new custom error with the provided message.
func NewCustomError(message string, details interface{}) *customError {
	return &customError{Message: message, Details: details}
}

// NewErrorFromContext creates a new custom error using the given context,
// and template string for multilanguage support.
func NewErrorFromContext(ctx ErrorCustomiser, template string, args ...interface{}) *customError {
	msg, err := ctx.CustomiseError(template, args...)
	if err != nil {
		return NewCustomError(err.Error(), nil)
	}
	return NewCustomError(msg, nil)
}

// DefaultErrorCustomiser implements ErrorCustomiser and provides basic multilanguage support.
type DefaultErrorCustomiser struct {
	Language string
}

// CustomiseError implements ErrorCustomiser's method.
func (c *DefaultErrorCustomiser) CustomiseError(template string, args ...interface{}) (string, error) {
	// Define multilanguage error message templates here
	var errorMessages = map[string]map[string]string{
		"en": {
			"generic":         "An error occurred: %s",
			"invalid_request": "Invalid request: %s",
			"not_found":       "Resource not found",
		},
		"es": {
			"generic":         "Ocurrió un error: %s",
			"invalid_request": "Solicitud inválida: %s",
			"not_found":       "Recurso no encontrado",
		},
	}

	lang := c.Language
	if lang == "" {
		lang = "en" // Fallback to English if no language is specified
	}

	msg, ok := errorMessages[lang][template]
	if !ok {
		return "", fmt.Errorf("unknown error message template: %s", template)
	}

	return fmt.Sprintf(msg, args...), nil
}

// RequestHandler serves HTTP requests and handles errors with custom error messages.
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user context from the request or set a default context as needed.
	// In this example, we'll use a default language.
	ctx := DefaultErrorCustomiser{Language: "en"}

	// ... business logic ...

	// Generate an error with customised message
	if err := doSomeProcessing(ctx); err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Request processed successfully"}`)
}

func doSomeProcessing(ctx ErrorCustomiser) error {
	// Simulate an error condition
	return NewErrorFromContext(ctx, "invalid_request", "missing required parameter")
}

func handleError(w http.ResponseWriter, err error) {
	if err, ok := err.(*customError); ok {
		// Handle custom errors with specific codes if needed
		http.Error(w, err.Message, http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		// Handle unknown errors generically
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewCustomError("An unknown error occurred", nil))
	}
}

func main() {
	http.HandleFunc("/", RequestHandler)
	http.ListenAndServe(":8080", nil)
}

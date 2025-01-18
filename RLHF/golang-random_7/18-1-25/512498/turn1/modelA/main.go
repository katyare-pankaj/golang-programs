package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// CustomErrorType is a definition of a structured error type.
type CustomErrorType struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Define some well-known error types
var (
	ErrNotFound     = &CustomErrorType{Code: "not_found", Message: "Resource not found"}
	ErrInvalidInput = &CustomErrorType{Code: "invalid_input", Message: "Invalid input"}
	ErrInternal     = &CustomErrorType{Code: "internal_error", Message: "Internal server error"}
)

// mapErrorToHTTPStatus maps custom error types to HTTP response statuses
func mapErrorToHTTPStatus(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrInvalidInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// ErrorMiddleware is a middleware that handles errors consistently across requests
func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				handleError(err, w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// handleError determines the error type and writes the appropriate HTTP response.
func handleError(err interface{}, w http.ResponseWriter) {
	var customErr *CustomErrorType

	// Handle different types of errors
	switch e := err.(type) {
	case *CustomErrorType:
		customErr = e
	case error:
		if errors.Is(e, ErrNotFound) {
			customErr = ErrNotFound
		} else if errors.Is(e, ErrInvalidInput) {
			customErr = ErrInvalidInput
		} else {
			customErr = ErrInternal
		}
	default:
		customErr = ErrInternal
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(mapErrorToHTTPStatus(customErr))
	_ = json.NewEncoder(w).Encode(customErr)
}

// ExampleHandler demonstrates a handler with potential for error
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate error to demonstrate the error handling
	if r.URL.Path != "/expected" {
		panic(ErrNotFound)
	}
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ExampleHandler)
	http.ListenAndServe(":8080", ErrorMiddleware(mux))
}

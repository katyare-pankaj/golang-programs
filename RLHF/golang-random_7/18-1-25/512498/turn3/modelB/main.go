package main

import (
	"fmt"
	"net/http"
)

// ErrorCode is an enum representing common error codes
type ErrorCode int

const (
	ErrorCodeInvalidRequest ErrorCode = iota + 1
	ErrorCodeNotFound
	ErrorCodeInternalServerError
	ErrorCodeUnauthorized
	// Add more error codes as needed
)

// AppError represents a custom application error
type AppError struct {
	Code    ErrorCode
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", e.Code, e.Message)
}

// ErrorResponse represents the response body for errors
type ErrorResponse struct {
	Error string `json:"error"`
}

// handleAppError sends an appropriate HTTP response based on the given error
func handleAppError(w http.ResponseWriter, err error, statusCode ...int) {
	appErr, ok := err.(*AppError)
	if !ok {
		appErr = &AppError{Code: ErrorCodeInternalServerError, Message: "Internal Server Error"}
	}

	w.Header().Set("Content-Type", "application/json")

	code := http.StatusInternalServerError // Default status code
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	switch appErr.Code {
	case ErrorCodeInvalidRequest:
		code = http.StatusBadRequest
	case ErrorCodeNotFound:
		code = http.StatusNotFound
	case ErrorCodeUnauthorized:
		code = http.StatusUnauthorized
	}

	w.WriteHeader(code)
	response := ErrorResponse{Error: appErr.Message}
	// You can also encode the response to JSON using encoding/json package if needed.
	fmt.Fprintf(w, "{\"error\":\"%s\"}", appErr.Message)
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleAppError(w, &AppError{Code: ErrorCodeInvalidRequest, Message: "Invalid method. Expected POST"})
		return
	}

	// Simulate some other operation that might return an error
	err := someOperation()
	if err != nil {
		handleAppError(w, err)
		return
	}

	w.Write([]byte("Operation succeeded!"))
}

func someOperation() error {
	// Simulate an error condition
	return &AppError{Code: ErrorCodeNotFound, Message: "Resource not found"}
}

func main() {
	http.HandleFunc("/example", exampleHandler)
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

// Define a custom error type
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Centralized error handling function
func handleAppError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*AppError); ok {
		http.Error(w, appErr.Message, appErr.Code)
		return
	}

	// Default to internal server error for unrecognized errors
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Example request handler
func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate an error
	err := someOperation()
	if err != nil {
		handleAppError(w, err)
		return
	}

	w.Write([]byte("Operation succeeded!"))
}

// Example function that returns a custom error
func someOperation() error {
	// Simulate an error condition
	return &AppError{Code: http.StatusBadRequest, Message: "Invalid operation"}
}

func main() {
	http.HandleFunc("/example", exampleHandler)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

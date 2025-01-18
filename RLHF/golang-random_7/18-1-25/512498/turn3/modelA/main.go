package main

import (
	"fmt"
	"net/http"
)

// Custom error struct
type AppError struct {
	Code    int
	Message string
}

// Implement the error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Error handling function
func handleAppError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*AppError); ok {
		http.Error(w, appErr.Message, appErr.Code)
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// Example HTTP handler using the custom error handling
func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate some operations which may produce an error
	if err := riskyOperation(); err != nil {
		handleAppError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Operation succeeded!"))
}

// Simulated operation that returns a custom error
func riskyOperation() error {
	return &AppError{Code: http.StatusBadRequest, Message: "Something went wrong with the request"}
}

func main() {
	http.HandleFunc("/example", exampleHandler)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

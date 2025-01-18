package main

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %d, message: %s", e.Code, e.Message)
}

func main() {
	// Define your custom error types with unique codes
	var (
		ErrNotFound            = &Error{Code: http.StatusNotFound, Message: "resource not found"}
		ErrInternalServerError = &Error{Code: http.StatusInternalServerError, Message: "internal server error"}
		ErrInvalidRequest      = &Error{Code: http.StatusBadRequest, Message: "invalid request"}
	)

	// Use the custom error types in your microservices
	myService := func(w http.ResponseWriter, r *http.Request) error {
		// Simulate an internal server error
		if r.URL.Path == "/error" {
			return ErrInternalServerError
		}

		switch r.Method {
		case http.MethodGet:
			// Handle GET request
			return nil
		case http.MethodPost:
			// Handle POST request
			return ErrInvalidRequest
		default:
			// Use ErrNotFound for unsupported methods
			return ErrNotFound
		}
	}

	// Wrap the microservice function with an error handler
	http.HandleFunc("/my-service", func(w http.ResponseWriter, r *http.Request) {
		err := myService(w, r)
		if err != nil {
			handleError(w, err)
		}
	})

	// Handle a route that simulates an internal server error
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		err := myService(w, r)
		if err != nil {
			handleError(w, err)
		}
	})

	fmt.Println("Server is running...")
	_ = http.ListenAndServe(":8080", nil)
}

func handleError(w http.ResponseWriter, err error) {
	// Implement your error handling logic here
	// You can extract the error code and message from the custom Error type
	if e, ok := err.(*Error); ok {
		http.Error(w, e.Message, e.Code)
		return
	}

	// Handle other errors (e.g., unexpected errors)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

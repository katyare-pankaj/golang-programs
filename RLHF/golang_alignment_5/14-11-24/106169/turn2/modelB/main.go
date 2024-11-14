package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// AppError defines a custom error for our application
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the string representation of the AppError
func (ae *AppError) Error() string {
	return ae.Message
}

// RequestLoggerMiddleware logs the request details
func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			latency := time.Since(start)
			log.Printf("%s %s %s %s %s",
				r.Method,
				r.URL,
				r.RemoteAddr,
				w.Header().Get("Content-Length"),
				latency.String(),
			)
		}()
		next.ServeHTTP(w, r)
	})
}

// TokenValidationMiddleware validates JWT tokens in the request header
func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// In a real application, you would validate the token here

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

// Example API endpoint
type Response struct {
	Message string `json:"message"`
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Hello, this is a protected API endpoint!"}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	r := mux.NewRouter()

	// Install middleware
	r.Use(RequestLoggerMiddleware)
	r.Use(TokenValidationMiddleware)

	// Protect your API endpoints with the TokenValidationMiddleware
	r.HandleFunc("/api/example", exampleHandler).Methods("GET")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

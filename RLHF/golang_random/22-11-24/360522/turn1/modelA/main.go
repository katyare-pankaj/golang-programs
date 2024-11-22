package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fatih/limiter"
	"github.com/gorilla/mux"
)

func getUserFromHeader(r *http.Request) string {
	return r.Header.Get("User-Id")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	userID := getUserFromHeader(r)

	// Create a limiter for 10 requests per minute per user
	limiterPerUser := limiter.NewRateLimiter(10, time.Minute)

	// Check if the request is allowed
	if !limiterPerUser.Allow(context.Background(), getUserFromHeader(r)) {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "Rate limit exceeded: %s", userID)
		return
	}

	// Handle the request
	fmt.Fprintf(w, "Hello, %s! Your request was allowed.", userID)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/data", handleRequest)

	log.Println("API server started on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type response struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

var cache map[string]response
var cacheMutex sync.Mutex
var defaultCacheExpirationTime = time.Duration(2) * time.Minute // Default cache expiration time if not specified by the user

func init() {
	cache = make(map[string]response)
}

func getPersonalizedResponse(name string) response {
	return response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Name:    name,
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the URL query parameters
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	values := u.Query()
	name := values.Get("name")
	expirationStr := values.Get("expiration")

	if name == "" {
		http.Error(w, "Name parameter required", http.StatusBadRequest)
		return
	}

	// Parse the expiration time from the query parameter (if provided)
	var cacheExpirationTime time.Duration
	if expirationStr != "" {
		expiration, err := strconv.Atoi(expirationStr)
		if err != nil {
			http.Error(w, "Invalid expiration time", http.StatusBadRequest)
			return
		}
		cacheExpirationTime = time.Duration(expiration) * time.Minute
	} else {
		cacheExpirationTime = defaultCacheExpirationTime
	}
	// Rest of the handling code remains the same...
}
func main() {
	http.HandleFunc("/personalize", handleRequest)
	fmt.Println("Server is running on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

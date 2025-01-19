package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type response struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

var cache map[string]response
var cacheMutex sync.Mutex
var cacheExpirationTime = time.Duration(2) * time.Minute // Cache expires after 2 minutes

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

	if name == "" {
		http.Error(w, "Name parameter required", http.StatusBadRequest)
		return
	}

	// Check if response is in cache
	cacheKey := name
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	cachedResponse, ok := cache[cacheKey]

	// Validate cache
	if ok {
		cacheExpiration := time.Now().Add(-cacheExpirationTime)
		if cachedResponse.Message != "" && cachedResponse.Name == name && time.Now().Before(cacheExpiration) {
			// Return cached response if valid
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(cachedResponse); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	// Generate new response
	newResponse := getPersonalizedResponse(name)
	cache[cacheKey] = newResponse

	// Return new response and cache it
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/personalize", handleRequest)
	fmt.Println("Server is running on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

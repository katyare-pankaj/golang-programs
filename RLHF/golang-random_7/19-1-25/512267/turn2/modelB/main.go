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
	Message string    `json:"message"`
	Name    string    `json:"name"`
	Expires time.Time `json:"expires"`
}

var cache map[string]response
var cacheMutex sync.Mutex
var cacheExpirationTime = time.Duration(2) * time.Minute // Cache expires after 2 minutes

func init() {
	cache = make(map[string]response)
	go cacheRefresher()
}

func getPersonalizedResponse(name string) response {
	return response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Name:    name,
		Expires: time.Now().Add(cacheExpirationTime),
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

	// Check if response is in cache and not expired
	cacheKey := name
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	cachedResponse, ok := cache[cacheKey]
	if ok && cachedResponse.Expires.After(time.Now()) {
		// Return cached response if valid
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cachedResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Generate new response if not in cache or expired
	newResponse := getPersonalizedResponse(name)
	cache[cacheKey] = newResponse

	// Return new response and cache it
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func cacheRefresher() {
	ticker := time.NewTicker(1 * time.Minute) // Refresh cache every minute
	defer ticker.Stop()

	for range ticker.C {
		cacheMutex.Lock()
		for name, resp := range cache {
			if resp.Expires.Before(time.Now()) {
				delete(cache, name) // Remove expired entries from the cache
			}
		}
		cacheMutex.Unlock()
	}
}

func main() {
	http.HandleFunc("/personalize", handleRequest)
	fmt.Println("Server is running on http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

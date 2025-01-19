package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type cacheEntry struct {
	response []byte
	expires  time.Time
}

var cache = make(map[string]cacheEntry)

func handler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	// Check if the cache entry exists and is not expired
	if entry, ok := cache[username]; ok && entry.expires.After(time.Now()) {
		// Return the cached response
		w.Write(entry.response)
		return
	}

	personalizedResponse := fmt.Sprintf("Hello, %s!", username)
	responseBytes := []byte(personalizedResponse)

	// Store the response in the cache with an expiration time of 1 minute
	cache[username] = cacheEntry{response: responseBytes, expires: time.Now().Add(time.Minute)}

	w.Write(responseBytes)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

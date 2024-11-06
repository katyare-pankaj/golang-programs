package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Sample data for the visualization
var data = []struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}{
	{"2023-01-01", 100.0},
	{"2023-01-02", 120.0},
	// More data entries...
}

// Cache struct to store the cached data
type Cache struct {
	Data     []byte
	Expiry   time.Time
	Duration time.Duration
}

// cachedData holds the cached data
var cachedData Cache
var mu sync.Mutex

func main() {
	// Initialize router and routes
	r := mux.NewRouter()
	r.HandleFunc("/data.json", dataHandler)

	// Start the server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the data is cached and valid
	mu.Lock()
	defer mu.Unlock()

	if time.Now().After(cachedData.Expiry) {
		// Cache is expired, generate new data and cache it
		cachedData.Data, _ = json.Marshal(data)
		cachedData.Expiry = time.Now().Add(time.Minute * 5) // Cache for 5 minutes
	}

	// Set the content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the cached data to the response
	_, _ = w.Write(cachedData.Data)
}

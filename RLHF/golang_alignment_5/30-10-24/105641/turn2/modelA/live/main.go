package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Flight struct represents flight data
type Flight struct {
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
}

var flights = []Flight{
	{
		Departure: time.Now().Add(1 * time.Hour).Format(time.RFC3339),
		Arrival:   time.Now().Add(3 * time.Hour).Format(time.RFC3339),
	},
	{
		Departure: time.Now().Add(2 * time.Hour).Format(time.RFC3339),
		Arrival:   time.Now().Add(4 * time.Hour).Format(time.RFC3339),
	},
}

func flightsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

func main() {
	http.HandleFunc("/api/flights", flightsHandler)
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}

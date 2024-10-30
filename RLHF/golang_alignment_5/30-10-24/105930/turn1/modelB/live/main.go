package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// sample flights data
var flights = []Flight{
	{
		DepartureTime: time.Now().Add(1 * time.Hour),
		ArrivalTime:   time.Now().Add(3 * time.Hour),
		Origin:        "JFK",
		Destination:   "LAX",
	},
	{
		DepartureTime: time.Now().Add(2 * time.Hour),
		ArrivalTime:   time.Now().Add(4 * time.Hour),
		Origin:        "ORD",
		Destination:   "DFW",
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

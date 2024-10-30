package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Flight struct represents a flight
type Flight struct {
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
}

func fetchFlights(url string, c chan<- []Flight) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching flights: %v", err)
		c <- nil
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c <- nil
		return
	}

	var flights []Flight
	if err := json.Unmarshal(body, &flights); err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
		c <- nil
		return
	}

	// Send the fetched flights back to the main goroutine via the channel
	c <- flights
}

func main() {
	apiURL := "http://localhost:8080/api/flights" // Use the local simulated API

	flightDataChan := make(chan []Flight)

	// Start the asynchronous flight data fetch
	go fetchFlights(apiURL, flightDataChan)

	// Main application loop
	for {
		select {
		case flights := <-flightDataChan:
			if flights == nil {
				fmt.Println("Error fetching flights.")
				break
			}
			// Process the fetched flight data
			for _, flight := range flights {
				fmt.Printf("Flight: %s to %s\n", flight.Departure, flight.Arrival)
			}
			// Optionally break the loop after processing
			return

		default:
			// Add other main application tasks here
			time.Sleep(time.Second) // Simulate some work
			fmt.Println("App is running...")
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Flight struct represents a flight data
type Flight struct {
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
}

func fetchFlights(url string) ([]Flight, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var flights []Flight
	if err := json.Unmarshal(body, &flights); err != nil {
		return nil, err
	}

	return flights, nil
}

func main() {
	apiURL := "https://example.com/api/flights" // Replace this with the actual API URL

	// Fetch flights asynchronously using goroutines
	go func() {
		flights, err := fetchFlights(apiURL)
		if err != nil {
			log.Printf("Error fetching flights: %v", err)
			return
		}

		// Display the flights on the user interface (simulated using fmt.Println)
		for _, flight := range flights {
			fmt.Printf("Flight: %s to %s, Departure: %s, Arrival: %s\n",
				flight.Origin, flight.Destination,
				flight.DepartureTime.Format("2006-01-02 15:04"),
				flight.ArrivalTime.Format("2006-01-02 15:04"))
		}
	}()

	// Run the main application loop
	select {}
}

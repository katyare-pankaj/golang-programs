package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TransportationData holds information about a vehicle and its distance
type TransportationData struct {
	Vehicle  string
	Distance float64
}

func processTransportationData(data *TransportationData, wg *sync.WaitGroup) {
	// Simulate processing the data (e.g., calculating travel time)
	<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Calculating the travel time based on distance and speed (assuming constant speed)
	timeTaken := data.Distance / 50

	fmt.Printf("Vehicle: %s, Distance: %.2fkm, Travel Time: %.2f hours\n", data.Vehicle, data.Distance, timeTaken)

	// Signal to the WaitGroup that this goroutine is done
	wg.Done()
}

func main() {
	// Transportation data records representing multiple vehicles
	transportationData := []TransportationData{
		{"Car1", 120.0},
		{"Train", 350.0},
		{"Airplane", 8000.0},
		{"Bicycle", 25.0},
		{"Bus", 400.0},
	}

	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Tell the WaitGroup that we will wait for 'len(transportationData)' goroutines
	wg.Add(len(transportationData))

	// Create a goroutine for each transportation data record to process concurrently
	for _, data := range transportationData {
		go processTransportationData(&data, &wg)
	}

	// Wait for all goroutines to complete processing
	wg.Wait()

	fmt.Println("All data processed successfully!")
}

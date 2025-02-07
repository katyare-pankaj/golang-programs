package main

import (
	"fmt"
	"sync"
	"time"
)

// TransportationData holds the data for a single transportation record.
type TransportationData struct {
	ID       int
	Vehicle  string
	Distance int // in kilometers
}

// ProcessData simulates processing each transportation data record.
// In a real application, this could involve complex computations, database operations, etc.
func ProcessData(data TransportationData, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	// Simulate processing time with sleep
	time.Sleep(time.Millisecond * 100)

	// Example processing: Just printing the data
	fmt.Printf("Processed record ID %d: Vehicle %s traveled %d km\n", data.ID, data.Vehicle, data.Distance)
}

func main() {
	// Example data
	dataRecords := []TransportationData{
		{ID: 1, Vehicle: "Bus", Distance: 120},
		{ID: 2, Vehicle: "Bike", Distance: 15},
		{ID: 3, Vehicle: "Car", Distance: 200},
	}

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Add the number of records to process to the WaitGroup's counter
	wg.Add(len(dataRecords))

	// Launch a goroutine to process each transportation data record
	for _, record := range dataRecords {
		// Pass the WaitGroup by reference
		go ProcessData(record, &wg)
	}

	// Wait for all goroutines to finish processing
	wg.Wait()
	fmt.Println("All data records processed.")
}

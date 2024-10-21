package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define a function that reads sensor data asynchronously using CPS
func readSensorDataAsync(sensorID int, cont func(float64, error)) {
	go func() {
		// Simulate reading sensor data with a random delay
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		// Generate random sensor data
		data := rand.Float64() * 100
		cont(data, nil)
	}()
}

// CPS function to process sensor data and send it to a channel
func processSensorDataAsync(data float64, sensorID int, ch chan<- float64, cont func()) {
	// Simulate processing data
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	// Send the processed data to the channel
	ch <- data
	cont()
}

// Continuation-passing style function to handle multiple sensor readings and processing them asynchronously
func handleSensorReadingsAsync(sensorIDs []int, ch chan<- float64, cont func()) {
	var wg sync.WaitGroup

	for _, sensorID := range sensorIDs {
		wg.Add(1)
		readSensorDataAsync(sensorID, func(data float64, err error) {
			if err != nil {
				fmt.Println("Error reading sensor data:", err)
				wg.Done()
				return
			}

			processSensorDataAsync(data, sensorID, ch, func() {
				wg.Done()
			})
		})
	}

	wg.Wait()
	cont()
}

func main() {
	// Continuous loop to handle real-time sensor data
	sensorIDs := []int{1, 2, 3, 4, 5}
	dataCh := make(chan float64)

	for {
		handleSensorReadingsAsync(sensorIDs, dataCh, func() {
			// Loop back to read more sensor data
		})

		// Consume the processed data
		select {
		case data := <-dataCh:
			fmt.Printf("Processed Sensor Data: %.2f\n", data)
		default:
			// No data available, keep looping
		}
	}
}

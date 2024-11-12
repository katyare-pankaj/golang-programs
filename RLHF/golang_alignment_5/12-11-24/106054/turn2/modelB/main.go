package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to receive data from the data streams
	dataChan := make(chan string)

	// Start a goroutine to simulate receiving data from a stream
	go receiveData(dataChan)

	// Start another goroutine to process the received data
	go processData(dataChan)

	// Wait for both goroutines to finish
	time.Sleep(2 * time.Second)
}

// receiveData simulates receiving data from a stream and sends it over the channel
func receiveData(dataChan chan<- string) {
	for i := 0; i < 5; i++ {
		data := fmt.Sprintf("Data %d", i)
		fmt.Println("Received:", data)
		dataChan <- data
		time.Sleep(1 * time.Second)
	}
	close(dataChan) // Close the channel to signal the end of data
}

// processData receives data from the channel and processes it
func processData(dataChan <-chan string) {
	for data := range dataChan {
		fmt.Println("Processing:", data)
		// Perform processing here
	}
}

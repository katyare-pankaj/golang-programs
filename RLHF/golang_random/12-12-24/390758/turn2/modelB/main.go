package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel for receiving data
	dataChannel := make(chan int)

	// Simulate a situation where data might not be immediately available
	go func() {
		time.Sleep(3 * time.Second) // Simulate delay
		dataChannel <- 42           // Send data after delay
	}()

	// Set a timeout of 2 seconds
	timeout := time.After(2 * time.Second)

	select {
	case data := <-dataChannel:
		fmt.Println("Received data:", data)
	case <-timeout:
		fmt.Println("Timed out waiting for data.")
	}
}

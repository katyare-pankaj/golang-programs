package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simulated external system data
type ExternalData struct {
	Value int
}

// Simulate external system
func simulateExternalSystem(dataChannel chan<- ExternalData, stopChannel <-chan struct{}) {
	for {
		select {
		case <-stopChannel:
			fmt.Println("External system stopped.")
			return
		default:
			// Generate random data
			data := ExternalData{Value: rand.Intn(100)}
			dataChannel <- data
			fmt.Println("External system sent data:", data)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
		}
	}
}

// Process data from external system
func processData(dataChannel <-chan ExternalData, resultChannel chan<- int, stopChannel <-chan struct{}) {
	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				fmt.Println("Processing stopped.")
				return
			}
			// Process the data (in this example, we just square it)
			processedResult := data.Value * data.Value
			resultChannel <- processedResult
			fmt.Println("Data processed:", data, "->", processedResult)
		case <-stopChannel:
			fmt.Println("Processing stopped.")
			return
		}
	}
}

// Request processed data
func requestData(resultChannel <-chan int, stopChannel <-chan struct{}) {
	for {
		select {
		case result, ok := <-resultChannel:
			if !ok {
				fmt.Println("Data request stopped.")
				return
			}
			fmt.Println("Received processed result:", result)
		case <-stopChannel:
			fmt.Println("Data request stopped.")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create channels
	dataChannel := make(chan ExternalData)
	resultChannel := make(chan int)
	stopChannel := make(chan struct{})

	// Start Goroutines
	go simulateExternalSystem(dataChannel, stopChannel)
	go processData(dataChannel, resultChannel, stopChannel)
	go requestData(resultChannel, stopChannel)

	// Wait for user input to stop
	var input string
	fmt.Scanln(&input)

	// Stop all Goroutines
	close(stopChannel)
	close(dataChannel)
	close(resultChannel)
}

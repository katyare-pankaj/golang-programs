package main

import (
	"fmt"
	"time"
)

// Event represents a type of event
type Event string

const (
	// Define different types of events
	EventTypeMessage  Event = "message"
	EventTypeError    Event = "error"
	EventTypeShutdown Event = "shutdown"
)

func main() {
	// Create channels for different types of events
	messageChan := make(chan Event)
	errorChan := make(chan Event)
	shutdownChan := make(chan Event)

	// Start goroutines to simulate events
	go sendMessages(messageChan)
	go sendErrors(errorChan)
	go func() {
		time.Sleep(5 * time.Second) // Simulate a delayed shutdown
		shutdownChan <- EventTypeShutdown
	}()

	// Main loop to handle events
	for {
		select {
		case message := <-messageChan:
			handleMessage(message)
		case error := <-errorChan:
			handleError(error)
		case <-shutdownChan:
			handleShutdown()
			fmt.Println("Shutting down...")
			return
		default:
			time.Sleep(100 * time.Millisecond) // Poll for new events
		}
	}
}

// sendMessages simulates sending message events
func sendMessages(chan Event) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(i+1) * time.Second)
		messageChan <- EventTypeMessage
	}
}

// sendErrors simulates sending error events
func sendErrors(chan Event) {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Duration(i+2) * time.Second)
		errorChan <- EventTypeError
	}
}

// handleMessage handles a message event
func handleMessage(event Event) {
	fmt.Printf("Received message: %s\n", event)
}

// handleError handles an error event
func handleError(event Event) {
	fmt.Printf("Received error: %s\n", event)
}

// handleShutdown handles a shutdown event
func handleShutdown() {
}

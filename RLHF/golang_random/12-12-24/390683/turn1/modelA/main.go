package main

import (
	"fmt"
	"time"
)

// Event types
type Event struct {
	Message string
}

type ErrorEvent struct {
	Error string
}

func main() {
	// Channels for different types of events
	msgChannel := make(chan Event)
	errChannel := make(chan ErrorEvent)
	timeoutChannel := time.After(5 * time.Second) // A timeout channel

	go func() {
		// Simulate sending messages
		for i := 0; i < 3; i++ {
			msgChannel <- Event{Message: fmt.Sprintf("Message %d", i)}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		// Simulate sending errors
		for i := 0; i < 3; i++ {
			errChannel <- ErrorEvent{Error: fmt.Sprintf("Error %d", i)}
			time.Sleep(2 * time.Second)
		}
	}()

	// Handle events using select
	for i := 0; i < 6; i++ { // We expect up to 6 events (3 messages + 3 errors)
		select {
		case msg := <-msgChannel:
			fmt.Println("Received:", msg.Message)
		case err := <-errChannel:
			fmt.Println("Received Error:", err.Error)
		case <-timeoutChannel:
			fmt.Println("Timeout: No events received.")
			return // Exit the loop if we hit timeout
		}
	}
}

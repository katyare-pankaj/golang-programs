package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel for receiving messages
	messageChannel := make(chan string)

	// Start a goroutine to simulate sending a message after a delay
	go func() {
		// Simulate a delay before sending the message
		time.Sleep(3 * time.Second)
		messageChannel <- "Hello, World!"
	}()

	// Define a timeout duration
	timeout := 2 * time.Second

	// Use select to wait for a message or a timeout
	select {
	case msg := <-messageChannel:
		fmt.Println("Received:", msg)
	case <-time.After(timeout):
		fmt.Println("Timeout: No message received within", timeout)
	}
}

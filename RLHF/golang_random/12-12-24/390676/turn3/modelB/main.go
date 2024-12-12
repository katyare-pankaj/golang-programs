package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan<- string, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Worker", id, "shutting down.")
			return
		case msg := <-ch:
			fmt.Printf("Worker %d received: %s\n", id, msg)
			time.Sleep(time.Second * 2) // Simulate work
		default:
			time.Sleep(time.Millisecond * 100) // Poll
		}
	}
}

func main() {
	ch := make(chan string)
	done := make(chan struct{})

	// Start workers
	for i := 1; i <= 3; i++ {
		go worker(i, ch, done)
	}

	// Send messages
	for _, msg := range []string{"Hello", "World"} {
		select {
		case ch <- msg: // Send message
			fmt.Println("Sent:", msg)
		case <-time.After(time.Second * 3): // Timeout after 3 seconds
			fmt.Println("Failed to send:", msg, "after timeout.")
		}
	}

	// Shutdown workers
	close(done)
	close(ch)
}

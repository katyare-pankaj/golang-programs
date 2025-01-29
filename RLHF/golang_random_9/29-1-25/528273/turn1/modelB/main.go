package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	ID   int
	Data string
}

func processMessage(msg Message) {
	// Simulate processing time
	time.Sleep(time.Millisecond * 10)

	// Do some processing on the message
	fmt.Printf("Processing message %d: %s\n", msg.ID, msg.Data)
}

func worker(messages <-chan Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range messages {
		// Defer cleanup tasks
		defer func() {
			// Simulate cleanup task 1
			time.Sleep(time.Millisecond * 1)
			fmt.Printf("Cleanup task 1 for message %d\n", msg.ID)
		}()

		defer func() {
			// Simulate cleanup task 2
			time.Sleep(time.Millisecond * 1)
			fmt.Printf("Cleanup task 2 for message %d\n", msg.ID)
		}()

		processMessage(msg)
	}
}

func main() {
	const numWorkers = 10
	const numMessages = 100

	messages := make(chan Message, numMessages)
	var wg sync.WaitGroup

	// Start workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(messages, &wg)
	}

	// Generate messages
	for i := 0; i < numMessages; i++ {
		messages <- Message{
			ID:   i,
			Data: fmt.Sprintf("Message %d", i),
		}
	}

	// Close the message channel to signal the workers to stop
	close(messages)

	// Wait for workers to finish
	wg.Wait()
}

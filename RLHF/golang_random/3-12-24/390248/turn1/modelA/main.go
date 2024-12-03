package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Consumer goroutine
func consumer(channel chan int, numConsumers int) {
	defer wg.Done()
	for message := range channel {
		fmt.Printf("Consumer %d received: %d\n", numConsumers, message)
		time.Sleep(time.Second / 5) // Simulate work
	}
}

// Producer goroutine
func producer(channel chan int, numMessages int) {
	for i := 0; i < numMessages; i++ {
		channel <- i                // Send message
		time.Sleep(time.Second / 5) // Simulate work
	}
	close(channel) // Signal no more messages
}

func main() {
	const numMessages = 5
	const numConsumers = 3

	// Create a channel to send messages
	channel := make(chan int, numMessages)

	// Start consumer goroutines
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(channel, i)
	}

	// Start producer goroutine
	go producer(channel, numMessages)

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines finished.")
}

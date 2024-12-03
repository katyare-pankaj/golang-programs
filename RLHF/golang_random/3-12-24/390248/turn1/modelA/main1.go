package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Consumer goroutine
func consumer(channel chan int, cond *sync.Cond, numConsumers int) {
	defer wg.Done()
	for {
		cond.L.Lock()
		defer cond.L.Unlock()

		if len(channel) == 0 {
			cond.Wait() // Wait if there are no messages
		}

		msg, ok := <-channel
		if !ok {
			return // Exit if channel is closed
		}
		fmt.Printf("Consumer %d received: %d\n", numConsumers, msg)
		time.Sleep(time.Second / 5) // Simulate work
	}
}

// Producer goroutine
func producer(channel chan int, cond *sync.Cond, numMessages int) {
	for i := 0; i < numMessages; i++ {
		channel <- i                // Send message
		time.Sleep(time.Second / 5) // Simulate work
	}
	close(channel) // Signal no more messages

	cond.L.Lock()
	defer cond.L.Unlock()
	cond.Broadcast() // Wake all waiting consumers
}

func main() {
	const numMessages = 5
	const numConsumers = 3

	// Create a channel to send messages
	channel := make(chan int, numMessages)
	cond := &sync.Cond{L: &sync.Mutex{}} // Create condition variable

	// Start consumer goroutines
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(channel, cond, i)
	}

	// Start producer goroutine
	go producer(channel, cond, numMessages)

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines finished.")
}

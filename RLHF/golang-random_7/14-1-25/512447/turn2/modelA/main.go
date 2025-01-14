package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	numGoroutines = 10
	numMessages   = 10000
	latencyTime   = 50 * time.Millisecond // Simulate 50ms latency
)

type message struct {
	data string
}

func main() {
	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2) // Producers + Consumers

	// Create a channel for message passing
	msgChan := make(chan message, numMessages)

	// Start the producer goroutines
	for i := 0; i < numGoroutines; i++ {
		go producer(i, msgChan, &wg)
	}

	// Start the consumer goroutines
	for i := 0; i < numGoroutines; i++ {
		go consumer(msgChan, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channel to indicate that no more messages will be sent
	close(msgChan)

	// Output final memory usage
	fmt.Printf("Final memory allocation (bytes): %d\n", runtime.NumGoroutine())
}

func producer(id int, msgChan chan<- message, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numMessages; i++ {
		// Simulate network latency by sleeping for a random duration
		time.Sleep(latencyTime * time.Duration(rand.Intn(5)+1))

		// Create a new message and send it to the channel
		msg := message{data: fmt.Sprintf("Producer %d: Message %d", id, i)}
		msgChan <- msg

		// Analyze memory usage
		currentMemory := runtime.ReadMemStats()
		fmt.Printf("Producer %d after %d messages: Alloc %d bytes, Sys %d bytes, NumGC %d\n", id, i+1, currentMemory.Alloc, currentMemory.Sys, currentMemory.NumGC)
	}
}

func consumer(msgChan <-chan message, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range msgChan {
		// Process the received message
		_ = msg.data

		// Analyze memory usage
		currentMemory := runtime.ReadMemStats()
		fmt.Printf("Consumer: Alloc %d bytes, Sys %d bytes, NumGC %d\n", currentMemory.Alloc, currentMemory.Sys, currentMemory.NumGC)
	}
}

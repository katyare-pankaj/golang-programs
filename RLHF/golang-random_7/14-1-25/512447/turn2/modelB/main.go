package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numGoroutines = 100
	numMessages   = 1_000_000
	latencyTime   = 100 * time.Millisecond // Simulate 100ms latency
)

type message struct {
	data string
}

var totalMemoryUsage uint64

func main() {
	// Create a WaitGroup to wait for all goroutines to complete.
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	// Start timer to track execution time
	start := time.Now()
	// Create a channel for message passing with a buffer.
	msgChan := make(chan message, numGoroutines*10)
	// Launch the producer goroutines.
	for i := 0; i < numGoroutines; i++ {
		go producer(i, msgChan, &wg)
	}
	// Launch the consumer goroutines.
	for i := 0; i < numGoroutines; i++ {
		go consumer(msgChan, &wg)
	}
	// Wait for all goroutines to finish.
	wg.Wait()
	// Print total time taken
	fmt.Printf("Total time taken: %s\n", time.Since(start))
	// Print total memory usage.
	fmt.Printf("Total memory usage: %v bytes\n", totalMemoryUsage)
}

// producer function that sends messages to the channel.
func producer(id int, msgChan chan<- message, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numMessages/numGoroutines; i++ {
		// Simulate network latency by sleeping for a random duration.
		time.Sleep(latencyTime * time.Duration(rand.Intn(5)+1))
		// Generate a random message data.
		data := fmt.Sprintf("Producer %d: Message %d", id, i)
		// Create a new message struct, tracking the allocation size.
		msg := new(message)
		msg.data = data
		atomic.AddUint64(&totalMemoryUsage, uint64(len(data)))
		// Send the message to the channel.
		msgChan <- *msg
	}
}

// consumer function that receives messages from the channel.
func consumer(msgChan <-chan message, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range msgChan {
		// Simulate some work with a sleep.
		time.Sleep(time.Millisecond * 10)
		// Process the received message.
		_ = msg.data
		// Check if the message struct needs to be cleaned up manually.
	}
}

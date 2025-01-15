package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numGoroutines = 100
	numMessages   = 1_000_000
	bufferSize    = 10
)

type message struct {
	data string
}

var totalMemoryUsage uint64
var msgChan chan message
var producerWg sync.WaitGroup
var consumerWg sync.WaitGroup

func init() {
	// Initialize the message channel with a buffer size
	msgChan = make(chan message, bufferSize)
}

func main() {
	// Start the timer to track execution time
	start := time.Now()

	// Launch the producer goroutines
	for i := 0; i < numGoroutines; i++ {
		producerWg.Add(1)
		go producer(i)
	}

	// Launch the consumer goroutines
	for i := 0; i < numGoroutines; i++ {
		consumerWg.Add(1)
		go consumer(i)
	}

	// Wait for all producers to finish
	producerWg.Wait()

	// Signal all consumers by sending a special message for each
	for i := 0; i < numGoroutines; i++ {
		msgChan <- message{data: "done"}
	}

	// Wait for all consumer goroutines to finish
	consumerWg.Wait()

	// Print total time taken
	fmt.Printf("Total time taken: %s\n", time.Since(start))

	// Print total memory usage
	fmt.Printf("Total memory usage: %v bytes\n", totalMemoryUsage)
}

func producer(id int) {
	defer producerWg.Done()
	for i := 0; i < numMessages/numGoroutines; i++ {

		// Generate a random message data
		data := fmt.Sprintf("Producer %d: Message %d", id, i)

		// Create a new message struct, tracking the allocation size
		msg := new(message)
		msg.data = data
		atomic.AddUint64(&totalMemoryUsage, uint64(len(data)))

		// Send the message to the channel
		msgChan <- *msg
	}
}

func consumer(id int) {
	defer consumerWg.Done()
	for {
		msg := <-msgChan
		if msg.data == "done" {
			// Stop consuming when the "done" signal is received
			return
		}

		// Process the received message
		_ = msg.data
	}
}

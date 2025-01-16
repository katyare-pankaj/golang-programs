package main

import (
	"fmt"
<<<<<<< HEAD
	"math/rand"
=======
>>>>>>> 82712e93aac8380b9aa86d34cef523dbc7c4af75
	"sync"
	"sync/atomic"
	"time"
)

const (
	numGoroutines = 100
	numMessages   = 1_000_000
<<<<<<< HEAD
	latencyTime   = 100 * time.Millisecond // Simulate 100ms latency
	bufferSize    = 100
=======
	bufferSize    = 10
>>>>>>> 82712e93aac8380b9aa86d34cef523dbc7c4af75
)

type message struct {
	data string
}

var totalMemoryUsage uint64
var msgChan chan message
<<<<<<< HEAD
var doneChan chan struct{}
var wg sync.WaitGroup
=======
var producerWg sync.WaitGroup
var consumerWg sync.WaitGroup
>>>>>>> 82712e93aac8380b9aa86d34cef523dbc7c4af75

func init() {
	// Initialize the message channel with a buffer size
	msgChan = make(chan message, bufferSize)
<<<<<<< HEAD
	doneChan = make(chan struct{})
=======
>>>>>>> 82712e93aac8380b9aa86d34cef523dbc7c4af75
}

func main() {
	// Start the timer to track execution time
	start := time.Now()

	// Launch the producer goroutines
	for i := 0; i < numGoroutines; i++ {
<<<<<<< HEAD
		wg.Add(1)
=======
		producerWg.Add(1)
>>>>>>> 82712e93aac8380b9aa86d34cef523dbc7c4af75
		go producer(i)
	}

	// Launch the consumer goroutines
	for i := 0; i < numGoroutines; i++ {
<<<<<<< HEAD
		wg.Add(1)
		go consumer(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(doneChan)
=======
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
>>>>>>> 82712e93aac8380b9aa86d34cef523dbc7c4af75

	// Print total time taken
	fmt.Printf("Total time taken: %s\n", time.Since(start))

	// Print total memory usage
	fmt.Printf("Total memory usage: %v bytes\n", totalMemoryUsage)
}

func producer(id int) {
<<<<<<< HEAD
	defer wg.Done()
	for i := 0; i < numMessages/numGoroutines; i++ {
		// Simulate network latency by sleeping for a random duration
		time.Sleep(latencyTime * time.Duration(rand.Intn(5)+1))
=======
	defer producerWg.Done()
	for i := 0; i < numMessages/numGoroutines; i++ {
>>>>>>> 82712e93aac8380b9aa86d34cef523dbc7c4af75

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
<<<<<<< HEAD
	defer wg.Done()
	for {
		select {
		case msg := <-msgChan:
			// Simulate some work with a sleep
			time.Sleep(time.Millisecond * 10)

			// Process the received message
			_ = msg.data
		case <-doneChan:
			// Check for the done signal from the main goroutine
			return
		}
=======
	defer consumerWg.Done()
	for {
		msg := <-msgChan
		if msg.data == "done" {
			// Stop consuming when the "done" signal is received
			return
		}

		// Process the received message
		_ = msg.data
>>>>>>> 82712e93aac8380b9aa86d34cef523dbc7c4af75
	}
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numGoroutines = 10
	numMessages   = 100
	latencyTime   = 100 * time.Millisecond // Simulate 100ms latency
	bufferSize    = 10
)

type message struct {
	data string
}

var totalMemoryUsage uint64
var msgChan chan message
var wg sync.WaitGroup

func init() {
	// Initialize the message channel with a buffer size
	msgChan = make(chan message, bufferSize)
}

func main() {
	start := time.Now()

	// Launch producer goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go producer(i)
	}

	// Launch consumer goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go consumer(i)
	}

	// Wait for all producers to finish before closing the channel
	wg.Wait()
	close(msgChan)

	// Wait for all consumers to finish processing after channel is closed
	wg.Wait()

	// Print execution time and memory usage
	fmt.Printf("Total time taken: %s\n", time.Since(start))
	fmt.Printf("Total memory usage: %v bytes\n", totalMemoryUsage)
}

func producer(id int) {
	defer wg.Done()
	for i := 0; i < numMessages/numGoroutines; i++ {
		// Simulate network latency by sleeping for a random duration
		time.Sleep(latencyTime * time.Duration(rand.Intn(5)+1))

		// Generate a random message data
		data := fmt.Sprintf("Producer %d: Message %d", id, i)
		msg := message{data: data}

		// Track memory usage
		atomic.AddUint64(&totalMemoryUsage, uint64(len(data)))

		// Send the message to the channel
		msgChan <- msg
		fmt.Printf("Producer %d sent: %s\n", id, msg.data)
	}
}

func consumer(id int) {
	defer wg.Done()
	for msg := range msgChan {
		// Simulate some work with a sleep
		time.Sleep(time.Millisecond * 10)

		// Process the received message
		fmt.Printf("Consumer %d received: %s\n", id, msg.data)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numGoroutines = 5                      // Start with a smaller number of goroutines
	numMessages   = 20                     // Limit the number of messages for observation
	latencyTime   = 100 * time.Millisecond // Simulate 100ms latency
)

type message struct {
	data string
}

func main() {
	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numGoroutines) // Add the number of producers

	// Create a channel for message passing
	msgChan := make(chan message, numMessages)

	// Start the producer goroutines
	for i := 0; i < numGoroutines; i++ {
		go producer(i, msgChan, &wg)
	}

	// Start the consumer goroutines
	var consumerWg sync.WaitGroup
	consumerWg.Add(numGoroutines) // Add the number of consumers
	for i := 0; i < numGoroutines; i++ {
		go consumer(msgChan, &consumerWg)
	}

	// Wait for all producers to finish
	wg.Wait()

	// After all producers are done, close the channel
	close(msgChan)

	// Wait for all consumers to finish
	consumerWg.Wait()
}

func producer(id int, msgChan chan<- message, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numMessages; i++ {
		// Simulate network latency by sleeping for a random duration
		time.Sleep(latencyTime * time.Duration(rand.Intn(5)+1))

		// Create a new message and send it to the channel
		msg := message{data: fmt.Sprintf("Producer %d: Message %d", id, i)}
		msgChan <- msg
	}
}

func consumer(msgChan <-chan message, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range msgChan {
		// Print the received message to ensure consumers are working
		fmt.Println("Consumed:", msg.data)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numGoroutines = 100
	numMessages   = 10000
	latencyTime   = 100 * time.Millisecond // Simulate 100ms latency
)

type message struct {
	data string
}

func main() {
	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

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

	// Analyze memory allocation using pprof
	// Note: You need to run this command in the terminal after running the go build command
	// go tool pprof -web memprofile.out
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
		// Process the received message
		_ = msg.data
	}
}

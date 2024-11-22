package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Constants for defining the number of producers, buffered channel size, and number of numbers to generate.
const (
	numProducers = 3
	channelSize  = 100
	numNumbers   = 1000
)

// Define the Message type to be sent over the channel.
type Message struct {
	Number int
	Source int // Indicates which producer sent the message.
}

func main() {
	// Initialize a buffered channel of size `channelSize`.
	messages := make(chan Message, channelSize)

	// Initialize a wait group for coordinating the goroutines.
	var wg sync.WaitGroup

	// Start producer goroutines.
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go produceMessages(&wg, messages, i)
	}

	// Start consumer goroutines.
	go consumeMessages(messages)

	// Wait for all goroutines to complete.
	wg.Wait()
}

// produceMessages function generates random numbers and sends them to the channel.
func produceMessages(wg *sync.WaitGroup, messages chan<- Message, producerID int) {
	defer wg.Done()
	for i := 0; i < numNumbers; i++ {
		number := rand.Intn(100) // Generate a random number between 0 and 99.
		messages <- Message{Number: number, Source: producerID}
		// Sleep for a small time to simulate work and to spread out production.
		time.Sleep(time.Millisecond * 10)
	}
}

// consumeMessages function receives and processes messages from the channel.
func consumeMessages(messages <-chan Message) {
	for msg := range messages {
		fmt.Printf("Received message: %d from producer %d\n", msg.Number, msg.Source)
		// Simulate work by sleeping for a small amount of time.
		time.Sleep(time.Millisecond * 10)
	}
}

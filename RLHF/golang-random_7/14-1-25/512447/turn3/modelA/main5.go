package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numProducers      = 10
	numConsumers      = 20
	numMessages       = 100000
	latencyMin        = 10 * time.Millisecond
	latencyMax        = 100 * time.Millisecond
	channelBufferSize = 10000
)

type message struct {
	data string
}

var totalMemoryUsage uint64

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Create a buffered channel for message passing
	msgChan := make(chan message, channelBufferSize)

	// Launch the producer goroutines
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(i, msgChan, &wg)
	}

	// Launch the consumer goroutines
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(msgChan, &wg)
	}

	// Close the channel after all producers are done
	go func() {
		wg.Wait()      // Wait for all producers and consumers to finish
		close(msgChan) // Close the channel after producers are done
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// Print total memory usage
	fmt.Printf("Total memory usage: %v bytes\n", totalMemoryUsage)
}

// producer function that sends messages to the channel
func producer(id int, msgChan chan<- message, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numMessages/numProducers; i++ {
		// Simulate network latency by sleeping for a random duration
		time.Sleep(time.Duration(rand.Intn(int(latencyMax-latencyMin)) + int(latencyMin)))

		// Generate a random message data
		data := fmt.Sprintf("Producer %d: Message %d", id, i)

		// Create a new message struct, tracking the allocation size
		msg := new(message)
		msg.data = data
		atomic.AddUint64(&totalMemoryUsage, uint64(len(data)))

		// Send the message to the channel
		msgChan <- *msg

		// Print memory usage after sending a certain number of messages
		if i%1000 == 0 {
			fmt.Printf("Producer %d: Sent %d messages, current memory usage: %v bytes\n", id, i, atomic.LoadUint64(&totalMemoryUsage))
		}
	}
}

// consumer function that receives messages from the channel
func consumer(msgChan <-chan message, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range msgChan {
		// Simulate some work with a sleep
		time.Sleep(time.Millisecond * 10)

		// Process the received message (dummy processing)
		_ = msg.data
	}
}

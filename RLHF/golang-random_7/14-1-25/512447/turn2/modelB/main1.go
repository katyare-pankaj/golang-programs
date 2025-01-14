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
	latencyTime   = 10 * time.Millisecond // Reduced latency
)

type message struct {
	data string
}

var totalMemoryUsage uint64

func main() {
	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	producerWg.Add(numGoroutines) // WaitGroup for producers
	consumerWg.Add(numGoroutines) // WaitGroup for consumers

	start := time.Now()
	msgChan := make(chan message, numGoroutines*10)

	// Start producers
	for i := 0; i < numGoroutines; i++ {
		go producer(i, msgChan, &producerWg)
	}

	// Close the channel after all producers are done
	go func() {
		producerWg.Wait() // Wait for producers
		close(msgChan)    // Close the channel
	}()

	// Start consumers
	for i := 0; i < numGoroutines; i++ {
		go consumer(msgChan, &consumerWg)
	}

	// Wait for all consumers
	consumerWg.Wait()

	fmt.Printf("Total time taken: %s\n", time.Since(start))
	fmt.Printf("Total memory usage: %v bytes\n", totalMemoryUsage)
}

func producer(id int, msgChan chan<- message, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numMessages/numGoroutines; i++ {
		time.Sleep(latencyTime * time.Duration(rand.Intn(3)+1)) // Simulate latency
		data := fmt.Sprintf("Producer %d: Message %d", id, i)
		msg := message{data: data}
		atomic.AddUint64(&totalMemoryUsage, uint64(len(data)))
		msgChan <- msg
	}
}

func consumer(msgChan <-chan message, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range msgChan { // Loop until channel is closed
		time.Sleep(time.Millisecond * 5) // Simulate work
		fmt.Println("Consumed:", msg.data)
	}
}

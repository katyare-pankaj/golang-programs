package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers    = 3  // Number of producer goroutines
	numConsumers    = 2  // Number of consumer goroutines
	bufferSize      = 10 // Size of the shared buffer
	productionItems = 15 // Total number of items to produce
)

var wg sync.WaitGroup

// producer generates random numbers and sends them to the consumer channel.
func producer(id int, out chan<- int) {
	defer wg.Done()
	for i := 0; i < productionItems/numProducers; i++ {
		item := rand.Intn(100)
		fmt.Printf("Producer %d produced item %d\n", id, item)
		out <- item
	}
}

// consumer receives items from the producer channel, processes them, and signals when complete.
func consumer(id int, in <-chan int, done chan bool) {
	defer wg.Done()
	for item := range in {
		fmt.Printf("Consumer %d is processing item %d\n", id, item)
		time.Sleep(time.Duration(item/10) * time.Microsecond) // Simulate processing time
	}
	done <- true
}

func main() {
	// Create a buffered channel to share items between producers and consumers
	items := make(chan int, bufferSize)

	// Create a channel to signal when consumers are done
	done := make(chan bool, numConsumers)

	wg.Add(numProducers + numConsumers)

	// Start producer goroutines
	for i := 0; i < numProducers; i++ {
		go producer(i+1, items)
	}

	// Start consumer goroutines
	for i := 0; i < numConsumers; i++ {
		go consumer(i+1, items, done)
	}

	// Wait for all producers to finish
	wg.Wait()
	close(items) // Close the channel to signal consumers that no more items will be produced

	// Wait for all consumers to finish
	for i := 0; i < numConsumers; i++ {
		<-done
	}

	fmt.Println("All consumers have finished processing")
}

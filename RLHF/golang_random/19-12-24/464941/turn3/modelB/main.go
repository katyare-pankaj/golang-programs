package main

import (
	"fmt"
	"sync"
)

const (
	numProducers = 10
	numConsumers = 5
	numItems     = 100
	bufferSize   = 10
)

func producer(id int, items chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numItems; i++ {
		items <- i
		fmt.Printf("Producer %d: Produced item %d\n", id, i)
	}
}

func consumer(id int, items <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range items {
		fmt.Printf("Consumer %d: Consumed item %d\n", id, item)
	}
}

func main() {
	items := make(chan int, bufferSize)
	var wg sync.WaitGroup

	// Start producers
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(i, items, &wg)
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(i, items, &wg)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()
	close(items)
}

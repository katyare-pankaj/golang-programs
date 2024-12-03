package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers = 3
	numConsumers = 2
	channelSize  = 10
)

func producer(id int, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		item := rand.Intn(100) + 1 // Generate a random number between 1 and 100
		fmt.Printf("Producer %d generated item %d\n", id, item)
		c <- item
		time.Sleep(time.Millisecond * 100) // Simulate work
	}
}

func consumer(id int, c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range c {
		fmt.Printf("Consumer %d consumed item %d\n", id, item)
		time.Sleep(time.Millisecond * 50) // Simulate work
	}
}

func main() {
	// Create a channel to communicate between producers and consumers
	c := make(chan int, channelSize)
	var wg sync.WaitGroup

	// Launch producers
	for p := 1; p <= numProducers; p++ {
		wg.Add(1)
		go producer(p, c, &wg)
	}

	// Launch consumers
	for c := 1; c <= numConsumers; c++ {
		wg.Add(1)
		go consumer(c, c, &wg)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()
	fmt.Println("All producers and consumers are done.")
}

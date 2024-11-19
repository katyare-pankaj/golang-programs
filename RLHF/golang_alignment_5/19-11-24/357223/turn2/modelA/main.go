package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		c <- i
		fmt.Println("Produced:", i)
		time.Sleep(time.Second)
	}
}

func consumer(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range c {
		fmt.Println("Consumed:", num)
		time.Sleep(time.Second)
	}
}

func main() {
	const numProducers = 2
	const numConsumers = 3

	// Create a channel for communication
	c := make(chan int)

	// Create a wait group to synchronize the main goroutine with producers and consumers
	var wg sync.WaitGroup
	wg.Add(numProducers)
	wg.Add(numConsumers)

	// Start producer goroutines
	for i := 0; i < numProducers; i++ {
		go producer(c, &wg)
	}

	// Start consumer goroutines
	for i := 0; i < numConsumers; i++ {
		go consumer(c, &wg)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()

	// Close the channel to signal consumers to stop
	close(c)
}

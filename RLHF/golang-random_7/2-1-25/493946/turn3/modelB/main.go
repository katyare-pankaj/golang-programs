package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define the type of data that will be processed by producers and consumers
type Data struct {
	Value int
}

// Producer function that generates data and sends it to the channel
func producer(id int, dataCh chan<- Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		data := Data{Value: rand.Intn(100)}
		fmt.Printf("Producer %d produced data: %d\n", id, data.Value)
		dataCh <- data
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

// Consumer function that receives data from the channel and processes it
func consumer(id int, dataCh <-chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		fmt.Printf("Consumer %d consumed data: %d\n", id, data.Value)
		// Simulate processing time
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	// Number of producers and consumers
	numProducers := 3
	numConsumers := 3

	// Create a buffered channel to hold data between producers and consumers
	dataCh := make(chan Data, 10)

	// Create a wait group to synchronize the main goroutine with producers and consumers
	var wg sync.WaitGroup

	// Start producers
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(i, dataCh, &wg)
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(i, dataCh, &wg)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()

	fmt.Println("All producers and consumers completed.")
}

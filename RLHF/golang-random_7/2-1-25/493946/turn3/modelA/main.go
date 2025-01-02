package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define the type of data being produced and consumed
type Data struct {
	ID    int
	Value string
}

func main() {
	// Create a buffered channel to hold data
	dataChannel := make(chan Data, 100)

	// Create a wait group to synchronize the main goroutine with the producers and consumers
	var wg sync.WaitGroup

	// Start multiple producers
	numProducers := 3
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(dataChannel, &wg)
	}

	// Start multiple consumers
	numConsumers := 2
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(dataChannel, &wg)
	}

	// Close the channel to signal that no more data will be produced
	close(dataChannel)

	// Wait for all producers and consumers to finish
	wg.Wait()

	fmt.Println("All data processed.")
}

// Producer function that generates data and sends it to the channel
func producer(dataChannel chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ { // Each producer generates 10 pieces of data
		data := Data{ID: i, Value: fmt.Sprintf("Producer %d: Data %d", i, i)}
		dataChannel <- data
		fmt.Printf("Producer %d sent data: %v\n", i, data)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simulate work
	}
}

// Consumer function that receives data from the channel and processes it
func consumer(dataChannel chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataChannel {
		fmt.Printf("Consumer received data: %v\n", data)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simulate work
	}
}

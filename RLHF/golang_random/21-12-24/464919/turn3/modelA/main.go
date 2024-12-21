package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Data is the type of data being produced and consumed
type Data struct {
	value int
}

// Producer produces random data and sends it to the channel
func Producer(ch chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Simulate producing data
		data := Data{value: rand.Intn(100)}
		ch <- data
		time.Sleep(time.Millisecond * 50) // Simulate production time
	}
}

// Consumer consumes data from the channel and prints it
func Consumer(ch chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		// Simulate consuming data
		time.Sleep(time.Millisecond * 100) // Simulate consumption time
		fmt.Printf("Consumed: %d\n", data.value)
	}
}

func main() {
	// Initialize a sync.WaitGroup to wait for all Goroutines to complete
	var wg sync.WaitGroup

	// Create a buffer channel to hold data between producers and consumers
	dataChannel := make(chan Data, 10)

	// Start 3 producers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Producer(dataChannel, &wg)
	}

	// Start 2 consumers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go Consumer(dataChannel, &wg)
	}

	// Simulate running the process for a while
	time.Sleep(5 * time.Second)

	// Close the channel to signal producers and consumers to stop
	close(dataChannel)

	// Wait for all producers and consumers to finish
	wg.Wait()

	fmt.Println("All producers and consumers have finished")
}

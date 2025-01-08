package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers  = 3
	numConsumers  = 5
	numItems      = 100
	producerSleep = 100 * time.Millisecond
	consumerSleep = 200 * time.Millisecond
)

// Producer generates random integers and sends them to the data channel.
func producer(id int, dataCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numItems; i++ {
		randInt := rand.Intn(1000)
		dataCh <- randInt
		fmt.Printf("Producer %d: Produced item %d\n", id, randInt)
		time.Sleep(producerSleep)
	}
	fmt.Printf("Producer %d: Exiting\n", id)
}

// Consumer receives random integers from the data channel and processes them.
func consumer(id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case randInt, ok := <-dataCh:
			if !ok {
				// dataCh has been closed, exit the loop
				fmt.Printf("Consumer %d: Data channel closed, exiting\n", id)
				return
			}
			fmt.Printf("Consumer %d: Consumed item %d\n", id, randInt)
			time.Sleep(consumerSleep)
		default:
			// No data available, do some other work or sleep
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	dataCh := make(chan int, 100) // Buffered channel to prevent blocking
	var wg sync.WaitGroup

	// Start producers
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(i+1, dataCh, &wg)
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(i+1, dataCh, &wg)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()

	// Close the data channel to indicate no more items will be produced
	close(dataCh)
	fmt.Println("All producers and consumers have finished.")
}

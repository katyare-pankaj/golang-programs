package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers = 3
	numConsumers = 3
	maxEvents    = 10
	sleepPeriod  = time.Millisecond * 200
)

var (
	wg      sync.WaitGroup
	msgChan chan int
)

// producer generates random integers and sends them via the message channel
func producer() {
	for i := 0; i < maxEvents; i++ {
		// Generate a random integer
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(100)

		// Send the random integer to the message channel
		select {
		case msgChan <- number:
			fmt.Printf("Producer %d: Sent integer %d\n", wg.Count(), number)
		default:
			fmt.Println("Producer: Channel is closed, stopping...")
			return
		}
		// Sleep for a random period
		time.Sleep(sleepPeriod)
	}
	wg.Done()
}

// consumer receives integers from the message channel and processes them
func consumer() {
	defer wg.Done()
	for {
		select {
		case number, ok := <-msgChan:
			if !ok {
				fmt.Println("Consumer: Channel is closed, stopping...")
				return
			}
			fmt.Printf("Consumer %d: Received integer %d, processed it\n", wg.Count(), number)
		default:
			time.Sleep(sleepPeriod)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize message channel
	msgChan = make(chan int)

	// Start producers
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer()
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer()
	}

	// Close the message channel after all producers are done
	go func() {
		wg.Wait()
		close(msgChan)
	}()

	// Wait for all producers and consumers to finish
	wg.Wait()

	fmt.Println("Producer-consumer simulation complete.")
}

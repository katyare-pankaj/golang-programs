package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer sends random numbers to the send-only channel.
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		num := rand.Intn(100) // Generate a random number
		fmt.Printf("Produced: %d\n", num)
		ch <- num                          // Send the number to the channel
		time.Sleep(time.Millisecond * 500) // Simulate work
	}
	close(ch) // Close the channel when done
}

// Consumer receives numbers from the receive-only channel.
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed. Consumer exiting.")
				return // Exit when the channel is closed
			}
			fmt.Printf("Consumed: %d\n", num)
		case <-time.After(time.Second): // Timeout to avoid blocking
			fmt.Println("No message received within 1 second.")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int) // Create a channel for integers

	var wg sync.WaitGroup
	wg.Add(2)

	go producer(ch, &wg) // Start the producer
	go consumer(ch, &wg) // Start the consumer

	wg.Wait() // Wait for both goroutines to finish

	fmt.Println("Main: All done.")
}

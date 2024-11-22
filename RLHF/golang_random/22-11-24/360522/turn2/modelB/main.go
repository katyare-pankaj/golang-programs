package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// numGenerator represents a function that generates a random number.
type numGenerator func() int

// producer generates random numbers and sends them to the given channel.
func producer(id int, numGen numGenerator, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		num := numGen()
		out <- num // Send the generated number to the channel
		fmt.Printf("Producer %d: Generated number %d\n", id, num)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Add some random delay
	}
}

// consumer consumes numbers from the channel and prints them.
func consumer(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fmt.Printf("Consumer: Consumed number %d\n", num)
	}
}

func main() {
	// Create a buffered channel with a capacity of 10
	out := make(chan int, 10)

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Define a function that generates a random number between 1 and 100
	numGen := func() int {
		return rand.Intn(100) + 1
	}

	// Start 3 producer goroutines
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go producer(i, numGen, out, &wg)
	}

	// Start a consumer goroutine
	wg.Add(1)
	go consumer(out, &wg)

	// Wait for all goroutines to complete
	wg.Wait()

	close(out) // Close the channel to signal the consumer that no more numbers will be sent
}

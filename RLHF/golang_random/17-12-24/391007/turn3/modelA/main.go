package main

import (
	"fmt"
	"sync"
)

// Produce generates numbers and sends them to a send-only channel
func produce(numbers chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		numbers <- i
		fmt.Printf("Produced: %d\n", i)
	}
	close(numbers) // Close the channel when done
}

// Consume receives numbers from a receive-only channel
func consume(numbers <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numbers { // Automatically handles channel closure
		fmt.Printf("Consumed: %d\n", num)
	}
}

func main() {
	var wg sync.WaitGroup
	numbers := make(chan int) // Create an unidirectional channel for integers

	wg.Add(2)
	go produce(numbers, &wg) // Start producer
	go consume(numbers, &wg) // Start consumer

	wg.Wait() // Wait for both goroutines to finish
}

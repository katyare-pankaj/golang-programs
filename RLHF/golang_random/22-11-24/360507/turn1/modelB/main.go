package main

import (
	"fmt"
	"sync"
)

// Producer function generates data and sends it to the channel.
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i // Send data to the channel
		fmt.Println("Produced:", i)
	}
	close(ch) // Close the channel when done producing
}

// Consumer function receives data from the channel and processes it.
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Println("Consumed:", data)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // Create a channel for integer data

	wg.Add(2) // Add 2 goroutines to the wait group
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("All done!")
}

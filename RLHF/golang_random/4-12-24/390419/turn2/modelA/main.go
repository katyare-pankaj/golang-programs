package main

import (
	"fmt"
	"sync"
	"time"
)

// producer function generates a series of numbers and sends them to the channel.
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i                           // Send number to the channel
		time.Sleep(time.Millisecond * 50) // Simulate work
	}
}

// consumer function receives numbers from the channel and prints them.
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
		time.Sleep(time.Millisecond * 200) // Simulate longer work
	}
}

func main() {
	// Create a buffered channel with a buffer size of 10
	ch := make(chan int, 10)
	wg := &sync.WaitGroup{}

	// Launch 10 producer goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, wg)
	}

	// Launch 1 consumer goroutine
	wg.Add(1)
	go consumer(ch, wg)

	// Wait for all goroutines to complete
	wg.Wait()
}

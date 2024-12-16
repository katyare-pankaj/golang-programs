package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numItems = 100
)

// Producer function generates numbers and sends them over the channel
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numItems; i++ {
		ch <- i
	}
}

// Consumer function receives numbers from the channel and processes them
func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ch {
		time.Sleep(10 * time.Millisecond) // Simulate time to process each item
	}
}

func measurePingPong(channelType string, bufferSize int) {
	start := time.Now()

	var wg sync.WaitGroup
	ch := make(chan int, bufferSize)

	wg.Add(1)
	go producer(1, ch, &wg)

	wg.Add(1)
	go consumer(1, ch, &wg)

	wg.Wait()
	close(ch)

	elapsed := time.Since(start)
	fmt.Printf("%s Channel with buffer size %d took %s to complete.\n", channelType, bufferSize, elapsed)
}

func main() {
	// Test with buffered channel of size 2
	measurePingPong("Buffered", 2)

	// Test with buffered channel of size 20
	measurePingPong("Buffered", 20)

	// Test with unbuffered channel
	measurePingPong("Unbuffered", 0)
}

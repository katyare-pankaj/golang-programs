package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer function that sends data to the channel
func producer(ch chan<- int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		value := rand.Intn(100)
		ch <- value
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10))) // simulate work
	}
}

// Consumer function that receives data from the channel
func consumer(ch <-chan int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		<-ch                              // Simulate processing
		time.Sleep(time.Millisecond * 20) // simulate work
	}
}

func measureTime(channelType string, bufferSize int, count int) {
	var wg sync.WaitGroup

	var ch interface{}
	if channelType == "buffered" {
		ch = make(chan int, bufferSize)
	} else {
		ch = make(chan int)
	}

	start := time.Now()

	wg.Add(1)
	go producer(ch.(chan int), count, &wg)

	wg.Add(1)
	go consumer(ch.(chan int), count, &wg)

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("%s Channel (buffer size = %d) took %s\n", channelType, bufferSize, elapsed)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const itemCount = 1000

	// Measure performance for unbuffered channel
	measureTime("Unbuffered", 0, itemCount)

	// Measure performance for buffered channel with a buffer size of 10
	measureTime("Buffered", 10, itemCount)

	// Measure performance for buffered channel with a larger buffer size of 100
	measureTime("Buffered", 100, itemCount)
}

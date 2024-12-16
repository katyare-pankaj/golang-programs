package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer function that sends data to the channel
func producer(ch chan<- int, count int) {
	for i := 0; i < count; i++ {
		value := rand.Intn(100)
		ch <- value
	}
	close(ch)
}

// Consumer function that receives data from the channel
func consumer(ch <-chan int, count int) {
	for i := 0; i < count; i++ {
		<-ch
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	const workCount = 100000
	var wg sync.WaitGroup

	fmt.Println("\nUnbuffered Channel Performance:")
	unbufferedStart := time.Now()

	// Unbuffered channel
	unbufferedCh := make(chan int)

	wg.Add(1)
	go producer(unbufferedCh, workCount)

	wg.Add(1)
	go consumer(unbufferedCh, workCount)

	wg.Wait()

	unbufferedEnd := time.Now()
	fmt.Printf("Unbuffered Channel Execution Time: %v\n", unbufferedEnd.Sub(unbufferedStart))

	fmt.Println("\nBuffered Channel Performance (buffer size = 1000):")
	bufferedStart := time.Now()

	// Buffered channel with a size of 1000
	bufferedCh := make(chan int, 1000)

	wg.Add(1)
	go producer(bufferedCh, workCount)

	wg.Add(1)
	go consumer(bufferedCh, workCount)

	wg.Wait()

	bufferedEnd := time.Now()
	fmt.Printf("Buffered Channel Execution Time: %v\n", bufferedEnd.Sub(bufferedStart))
}

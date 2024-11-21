package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for msg := range ch {
		fmt.Println("Consumed:", msg)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// Unbuffered channel example
	fmt.Println("Unbuffered channel example:")
	unbufferedCh := make(chan int)
	wg.Add(1)
	go consumer(unbufferedCh, &wg)
	producer(unbufferedCh, 10)
	wg.Wait()
	fmt.Println("Unbuffered channel finished.")

	// Buffered channel example
	fmt.Println("\nBuffered channel example:")
	bufferedCh := make(chan int, 5) // Buffer size of 5
	wg.Add(1)
	go consumer(bufferedCh, &wg)
	producer(bufferedCh, 10)
	wg.Wait()
	fmt.Println("Buffered channel finished.")
}

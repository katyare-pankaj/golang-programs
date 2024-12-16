package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Println("Consumed:", msg)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	// Unbuffered channel
	unbufferedCh := make(chan int)

	wg.Add(2)
	go producer(unbufferedCh, &wg)
	go consumer(unbufferedCh, &wg)

	wg.Wait()
	fmt.Println("Unbuffered channel done.")

	// Buffered channel with size 10
	bufferedCh := make(chan int, 10)

	wg.Add(2)
	go producer(bufferedCh, &wg)
	go consumer(bufferedCh, &wg)

	wg.Wait()
	fmt.Println("Buffered channel done.")
}

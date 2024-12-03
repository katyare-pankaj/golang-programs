package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for item := range ch {
		fmt.Println("Consumed:", item)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	// Wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Wait()
	}()
}

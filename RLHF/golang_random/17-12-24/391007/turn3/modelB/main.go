package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i // Only producers send on this channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch { // Only consumers receive from this channel
		fmt.Println("Consumed:", num)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)

	// Wait for both goroutines to finish
	time.Sleep(500 * time.Millisecond)
}

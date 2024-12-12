package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(1 * time.Second) // Simulate work
	}
	close(ch) // Close the channel when done
}

func consumer(ch <-chan int) {
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Consumer: Channel closed, exiting.")
				return
			}
			fmt.Printf("Consumed: %d\n", val)
		case <-time.After(2 * time.Second):
			fmt.Println("Consumer: Timeout after 2 seconds, exiting.")
			return
		}
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)
}

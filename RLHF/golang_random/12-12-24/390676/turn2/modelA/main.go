package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(time.Second) // Simulate work
	}
	close(ch) // Close the channel when done
}

func consumer(ch <-chan int) {
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed! Exiting consumer.")
				return
			}
			fmt.Printf("Consumed: %d\n", val)
		case <-time.After(2 * time.Second):
			fmt.Println("No data received in 2 seconds, timeout.")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

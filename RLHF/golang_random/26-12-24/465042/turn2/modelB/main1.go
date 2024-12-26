package main

import (
	"fmt"
	"time"
)

func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Println("Received:", i)
	}
}
func main() {
	ch := make(chan int)
	go consumer(ch)
	// Send data to the channel
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all the data
	// Allow goroutines to finish their tasks before exiting
	time.Sleep(1 * time.Second)
}

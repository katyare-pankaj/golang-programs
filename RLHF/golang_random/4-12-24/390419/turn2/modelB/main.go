package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channel example
	unbuffered := make(chan int)

	go func() {
		unbuffered <- 1 // This will block until the main goroutine receives.
		fmt.Println("Sent to unbuffered channel")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Received from unbuffered channel:", <-unbuffered)

	// Buffered channel example
	buffered := make(chan int, 2)

	buffered <- 1 // This does not block because the channel is not full.
	buffered <- 2 // This also does not block because the buffer has capacity.

	fmt.Println("Buffered channel has capacity, sending done.")

	go func() {
		time.Sleep(1 * time.Second)
		buffered <- 3 // This will block until a receive is performed.
		fmt.Println("Sent to buffered channel after delay")
	}()

	// Receiving values from the buffered channel
	fmt.Println("Received from buffered channel:", <-buffered)
	fmt.Println("Received from buffered channel:", <-buffered)
	fmt.Println("Received from buffered channel:", <-buffered)
}

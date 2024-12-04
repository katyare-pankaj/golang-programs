package main

import (
	"fmt"
	"time"
)

// worker simulates a goroutine that performs work and sends results through a channel.
func worker(id int, ch chan<- string) {
	time.Sleep(time.Second)
	ch <- fmt.Sprintf("worker %d completed", id)
}

func main() {
	// Create a channel to communicate between main and worker goroutines.
	ch := make(chan string)

	// Launch multiple worker goroutines.
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Receive and print results from the workers.
	for i := 1; i <= 3; i++ {
		message := <-ch // This receive operation will block until a worker sends a message.
		fmt.Println(message)
	}

	// Close the channel as we no longer need to send/receive messages.
	close(ch)
}
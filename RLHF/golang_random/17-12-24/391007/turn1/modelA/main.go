package main

import (
	"fmt"
	"time"
)

// Function that sends messages to a send-only channel.
func producer(ch chan<- string) {
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("message %d", i)
		ch <- msg
		time.Sleep(time.Second) // Simulate work
	}
	close(ch) // Close the channel when done
}

// Function that receives messages from a receive-only channel.
func consumer(ch <-chan string) {
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}

func main() {
	ch := make(chan string)

	go producer(ch)
	go consumer(ch)

	// Using select with a timeout to avoid blocking indefinitely
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed, exiting.")
				return
			}
			fmt.Println("Received in select:", msg)
		case <-time.After(3 * time.Second):
			fmt.Println("No message received within 3 seconds.")
		}
	}
}

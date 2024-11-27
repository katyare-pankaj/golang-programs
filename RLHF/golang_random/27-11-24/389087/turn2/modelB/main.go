package main

import (
	"fmt"
	"sync"
	"time"
)

// Define the message type
type Message struct {
	Data string
}

// Publisher function - sends updates to subscribers
func Publisher(ch chan<- Message) {
	for i := 1; i <= 5; i++ {
		msg := Message{Data: fmt.Sprintf("Update %d", i)}
		ch <- msg
		time.Sleep(time.Second)
	}
	close(ch)
}

// Subscriber function - reads updates from the channel
func Subscriber(name string, ch <-chan Message, wg *sync.WaitGroup) {
	for msg := range ch {
		fmt.Printf("Subscriber %s received: %s\n", name, msg.Data)
	}
	wg.Done()
}

func main() {
	// Create a buffered channel with a capacity of 100
	ch := make(chan Message, 100)
	var wg sync.WaitGroup

	// Start the publisher
	go Publisher(ch)

	// Start multiple subscribers
	numSubscribers := 3
	wg.Add(numSubscribers)
	for i := 0; i < numSubscribers; i++ {
		go Subscriber(fmt.Sprintf("Subscriber%d", i), ch, &wg)
	}

	// Wait for all subscribers to finish
	wg.Wait()
}

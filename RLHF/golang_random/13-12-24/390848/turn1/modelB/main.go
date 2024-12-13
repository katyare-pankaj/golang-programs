package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func processMessages(msgChan chan string) {
	defer closeChan(msgChan)

	wg := &sync.WaitGroup{}

	// Simulate a worker goroutine
	go func() {
		defer wg.Done()
		for msg := range msgChan {
			fmt.Println("Processing message:", msg)
		}
	}()

	wg.Add(1)
	defer wg.Done()

	select {
	case msg := <-msgChan:
		fmt.Println("Received message:", msg)
		// Simulate an error that causes the function to return early
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out waiting for messages.")
	}

	// Normally, we would send more messages here
}

func closeChan(chanName string, ch chan string) {
	if ch == nil {
		log.Printf("Error closing channel %s: channel is nil\n", chanName)
		return
	}

	err := ch.Close()
	if err != nil {
		log.Printf("Error closing channel %s: %v\n", chanName, err)
	}
}

func main() {
	msgChan := make(chan string)

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered: %v\n", r)
		}
	}()

	// Start the message processing
	processMessages(msgChan)

	// Send a message to the channel
	select {
	case msgChan <- "Hello, World!":
	case <-time.After(500 * time.Millisecond):
		log.Println("Failed to send message before channel closed.")
	}

	// Keep the program running for a short while to observe output
	time.Sleep(2 * time.Second)
}

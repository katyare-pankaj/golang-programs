package main

import (
	"fmt"
	"time"
)

func main() {
	// Start with a slice of channels
	channels := make([]chan string, 0)

	// Create channels at runtime
	channels = append(channels, make(chan string))
	channels = append(channels, make(chan string))
	channels = append(channels, make(chan string))

	go func() {
		time.Sleep(time.Second)
		channels[0] <- "Channel 1: Message received"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		channels[1] <- "Channel 2: Message received"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		channels[2] <- "Channel 3: Message received"
	}()

	// Listen for messages from the dynamic channels
	for {
		select {
		case message := <-channels[0]:
			fmt.Println(message)
		case message := <-channels[1]:
			fmt.Println(message)
		case message := <-channels[2]:
			fmt.Println(message)
		default:
			// Optionally add a timeout or other behavior here
			time.Sleep(time.Millisecond * 100)
		}
	}
}

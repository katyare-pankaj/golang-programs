package main

import (
	"fmt"
	"time"
)

// Define a ChannelBroadcaster
type ChannelBroadcaster struct {
	send      chan<- string
	closed    bool
	consumers []chan string
}

// NewChannelBroadcaster initializes a new ChannelBroadcaster
func NewChannelBroadcaster() *ChannelBroadcaster {
	c := &ChannelBroadcaster{
		send: make(chan<- string),
	}
	go c.consumeProcessor()
	return c
}

// Send sends a message to all connected consumers
func (c *ChannelBroadcaster) Send(msg string) {
	c.send <- msg
}

// Close closes the broadcaster
func (c *ChannelBroadcaster) Close() {
	close(c.send)
	c.closed = true
}

// Connect adds a new consumer to the broadcaster
func (c *ChannelBroadcaster) Connect() <-chan string {
	if c.closed {
		return nil
	}
	consumerChan := make(chan string)
	c.consumers = append(c.consumers, consumerChan)
	return consumerChan
}

// consumeProcessor continuously listens for messages
func (c *ChannelBroadcaster) consumeProcessor() {
	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				for _, consumer := range c.consumers {
					close(consumer)
				}
				return
			}
			// Send message to all consumers
			for _, consumer := range c.consumers {
				select {
				case consumer <- msg:
				case <-time.After(1 * time.Second):
					// Optional: Close the consumer if it is slow to receive
					close(consumer)
				}
			}
		}
	}
}

func main() {
	broadcaster := NewChannelBroadcaster()

	// Launch consumers
	go func() {
		for msg := range broadcaster.Connect() {
			fmt.Println("Consumer 1:", msg)
		}
		fmt.Println("Consumer 1 ended")
	}()

	go func() {
		for msg := range broadcaster.Connect() {
			fmt.Println("Consumer 2:", msg)
		}
		fmt.Println("Consumer 2 ended")
	}()

	// Broadcast messages
	broadcaster.Send("Hello!")
	time.Sleep(1 * time.Second)
	broadcaster.Send("World!")

	// Close the broadcaster
	time.Sleep(2 * time.Second)
	broadcaster.Close()
}

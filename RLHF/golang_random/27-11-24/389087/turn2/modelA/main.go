package main

import (
	"fmt"
	"sync"
	"time"
)

// Define the type of messages to be published
type Message struct {
	Content   string
	Timestamp time.Time
}

// Create a publish-subscribe interface
type PublisherSubscriber interface {
	Publish(msg Message)
	Subscribe() <-chan Message
	Close()
}

// Implement the publisher-subscriber interface
type psImpl struct {
	messages    chan Message
	subscribers sync.WaitGroup
	cancel      chan struct{}
}

func (ps *psImpl) Publish(msg Message) {
	ps.messages <- msg
}

func (ps *psImpl) Subscribe() <-chan Message {
	ps.subscribers.Add(1)
	defer ps.subscribers.Done()

	return ps.messages
}

func (ps *psImpl) Close() {
	close(ps.messages)
	ps.subscribers.Wait()
	close(ps.cancel)
}

func main() {
	// Create a new publisher-subscriber instance
	ps := &psImpl{
		messages:    make(chan Message, 100), // Buffered channel to avoid blocking
		subscribers: sync.WaitGroup{},
		cancel:      make(chan struct{}),
	}

	// Start a publisher goroutine
	go func() {
		defer close(ps.cancel)
		for {
			select {
			case <-ps.cancel:
				return
			default:
				msg := Message{
					Content:   fmt.Sprintf("Message %d", time.Now().UnixNano()),
					Timestamp: time.Now(),
				}
				ps.Publish(msg)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Start multiple subscriber goroutines
	numSubscribers := 3
	for i := 0; i < numSubscribers; i++ {
		go func(id int) {
			subscription := ps.Subscribe()
			for msg := range subscription {
				fmt.Printf("Subscriber %d received: %s at %s\n", id, msg.Content, msg.Timestamp)
			}
			fmt.Printf("Subscriber %d is done.\n", id)
		}(i)
	}

	// Simulate a graceful shutdown
	time.Sleep(5 * time.Second)
	ps.Close()
}

package main

import (
	"sync"
	"time"
)

type Actor struct {
	state     string // Can be "initialized", "processing", "terminated"
	waitGroup sync.WaitGroup
	messages  chan string
}

func NewActor() *Actor {
	return &Actor{
		state:     "initialized",
		waitGroup: sync.WaitGroup{},
		messages:  make(chan string, 10),
	}
}

func (a *Actor) Initialize() {
	a.waitGroup.Add(1)
	defer a.waitGroup.Done()

	a.state = "initialized"
	printf("Actor initialized\n")
}

func (a *Actor) Process() {
	a.waitGroup.Add(1)
	defer a.waitGroup.Done()

	a.state = "processing"
	printf("Actor started processing\n")

	for {
		select {
		case message := <-a.messages:
			printf("Received message: %s\n", message)
			time.Sleep(time.Second)
		case <-time.After(time.Second * 5):
			printf("Actor finished processing\n")
			return
		}
	}
}

func (a *Actor) Terminate() {
	a.waitGroup.Add(1)
	defer a.waitGroup.Done()

	close(a.messages)
	a.state = "terminated"
	printf("Actor terminated\n")
}

func (a *Actor) Wait() {
	a.waitGroup.Wait()
}

func main() {
	actor := NewActor()

	// Initialize the actor
	go actor.Initialize()

	// Process messages
	go actor.Process()

	// Wait for initialization to complete
	actor.waitGroup.Wait()

	// Send messages to the actor
	actor.messages <- "Hello"
	actor.messages <- "World"

	// Wait for processing to complete
	actor.waitGroup.Wait()

	// Terminate the actor
	go actor.Terminate()

	// Wait for termination to complete
	actor.waitGroup.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// Message is the interface for messages sent between actors.
type Message interface{}

// Actor is the interface for actors.
type Actor interface {
	Start()
	Stop()
	Send(Message)
}

// SimpleActor is a basic actor that processes messages and transitions between states.
type SimpleActor struct {
	name     string
	state    string
	inbox    chan Message
	wg       sync.WaitGroup
	quit     chan struct{}
	stopOnce sync.Once
}

// NewSimpleActor creates a new SimpleActor.
func NewSimpleActor(name string) *SimpleActor {
	return &SimpleActor{
		name:  name,
		state: "idle",
		inbox: make(chan Message),
		quit:  make(chan struct{}),
	}
}

// Start starts the actor's processing loop.
func (a *SimpleActor) Start() {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		for {
			select {
			case msg := <-a.inbox:
				a.processMessage(msg)
			case <-a.quit:
				return
			}
		}
	}()
}

// Stop stops the actor's processing loop.
func (a *SimpleActor) Stop() {
	a.stopOnce.Do(func() {
		close(a.quit)
	})
}

// Send sends a message to the actor's inbox.
func (a *SimpleActor) Send(msg Message) {
	a.inbox <- msg
}

// processMessage processes a message received by the actor.
func (a *SimpleActor) processMessage(msg Message) {
	switch msg := msg.(type) {
	case string:
		a.handleStringMessage(msg)
	default:
		fmt.Printf("Actor %s: Unrecognized message: %v\n", a.name, msg)
	}
}

// handleStringMessage handles a string message and transitions between states.
func (a *SimpleActor) handleStringMessage(msg string) {
	switch a.state {
	case "idle":
		fmt.Printf("Actor %s: Received message '%s' in idle state. Transitioning to processing state.\n", a.name, msg)
		a.state = "processing"
		go a.processLongRunningTask(msg)
	case "processing":
		fmt.Printf("Actor %s: Received message '%s' in processing state. Ignoring.\n", a.name, msg)
	default:
		fmt.Printf("Actor %s: Received message '%s' in unknown state '%s'. Ignoring.\n", a.name, msg, a.state)
	}
}

// processLongRunningTask simulates a long-running task and transitions back to the idle state when complete.
func (a *SimpleActor) processLongRunningTask(msg string) {
	defer func() {
		a.state = "idle"
		fmt.Printf("Actor %s: Long-running task for message '%s' completed. Transitioning back to idle state.\n", a.name, msg)
	}()

	fmt.Printf("Actor %s: Starting long-running task for message '%s'...\n", a.name, msg)
	// Simulate long-running task with time.Sleep
	time.Sleep(2 * time.Second) // Simulate a 2-second long task
	fmt.Printf("Actor %s: Long-running task for message '%s' finished.\n", a.name, msg)

	// Ensure the WaitGroup counter is decremented after the task is completed
	a.wg.Done()
}

// Wait waits for the actor to complete its processing loop.
func (a *SimpleActor) Wait() {
	a.wg.Wait()
}

func main() {
	// Create actors
	actor1 := NewSimpleActor("actor1")
	actor2 := NewSimpleActor("actor2")

	// Start actors
	actor1.Start()
	actor2.Start()

	// Send messages to actors
	actor1.Send("hello")
	actor1.Send("world")
	actor2.Send("foo")
	actor2.Send("bar")

	// Wait for actors to complete
	actor1.Wait()
	actor2.Wait()

	// Stop actors
	actor1.Stop()
	actor2.Stop()
}

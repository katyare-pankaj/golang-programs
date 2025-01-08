package main

import (
	"fmt"
	"sync"
	"time"
)

// Actor represents a basic actor with an inbox and a name.
type Actor struct {
	name     string
	inbox    chan Message
	state    string
	wg       sync.WaitGroup
	stopOnce sync.Once
}

// Message defines the interface for messages that actors can receive.
type Message interface {
	handle(Actor)
}

// StringMessage is a simple message that holds a string value.
type StringMessage string

// handle implements the Message interface for StringMessage.
func (msg StringMessage) handle(a Actor) {
	fmt.Printf("Actor %s received message: %s\n", a.name, msg)
}

// NewActor creates a new Actor with a specified name.
func NewActor(name string) *Actor {
	return &Actor{
		name:  name,
		inbox: make(chan Message, 100),
		state: "idle",
	}
}

// Start starts the actor's message processing loop.
func (a *Actor) Start() {
	a.wg.Add(1)
	defer a.wg.Done()

	for {
		select {
		case msg := <-a.inbox:
			if a.state == "idle" {
				msg.handle(a)
				a.state = "processing"
				go a.processTask(msg)
			} else {
				fmt.Printf("Actor %s is busy and cannot handle message: %s\n", a.name, msg)
			}
		case <-a.stop():
			return
		}
	}
}

// Stop stops the actor gracefully.
func (a *Actor) Stop() {
	a.stopOnce.Do(func() {
		close(a.inbox)
	})
}

// Wait waits for the actor to finish processing all messages.
func (a *Actor) Wait() {
	a.wg.Wait()
}

// stop returns a channel that is closed when the actor is stopped.
func (a *Actor) stop() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		a.Stop()
		close(ch)
	}()
	return ch
}

// processTask simulates a long-running task and transitions back to the idle state.
func (a *Actor) processTask(msg Message) {
	defer func() {
		a.state = "idle"
		fmt.Printf("Actor %s finished processing task: %s\n", a.name, msg)
	}()

	time.Sleep(2 * time.Second)
}

func main() {
	// Create actors
	actor1 := NewActor("Actor1")
	actor2 := NewActor("Actor2")

	// Start actors
	actor1.Start()
	actor2.Start()

	// Send messages to actors
	actor1.inbox <- StringMessage("Hello")
	actor1.inbox <- StringMessage("World")
	actor2.inbox <- StringMessage("Foo")
	actor2.inbox <- StringMessage("Bar")

	// Wait for actors to finish
	actor1.Wait()
	actor2.Wait()

	// Stop actors
	actor1.Stop()
	actor2.Stop()
}

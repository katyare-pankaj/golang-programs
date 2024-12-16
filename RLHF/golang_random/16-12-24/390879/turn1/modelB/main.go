package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// Event is the interface that all events must implement
type Event interface {
	Apply()
}

// AccountCreatedEvent represents an account creation event
type AccountCreatedEvent struct {
	AccountID string
}

func (e *AccountCreatedEvent) Apply() {
	fmt.Println("Account created:", e.AccountID)
}

// DepositEvent represents a deposit event
type DepositEvent struct {
	AccountID string
	Amount    float64
}

func (e *DepositEvent) Apply() {
	fmt.Println("Deposit made to", e.AccountID, ": $", e.Amount)
}

// EventHandler uses reflection to handle events
func EventHandler(events chan Event) {
	for event := range events {
		val := reflect.ValueOf(event)
		typ := val.Type()

		fmt.Printf("Handling event of type %s\n", typ.Name())

		// Call the Apply method using reflection
		val.MethodByName("Apply").Call(nil)
	}
}

func main() {
	// Create a channel for events
	events := make(chan Event)

	// Start a goroutine to handle events
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		EventHandler(events)
		wg.Done()
	}()

	// Simulate concurrent event generation
	time.AfterFunc(1*time.Second, func() {
		events <- &AccountCreatedEvent{AccountID: "123"}
	})
	time.AfterFunc(2*time.Second, func() {
		events <- &DepositEvent{AccountID: "123", Amount: 100.0}
	})
	time.AfterFunc(3*time.Second, func() {
		events <- &DepositEvent{AccountID: "123", Amount: 200.0}
		close(events)
	})

	wg.Wait()
}

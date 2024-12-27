package main

import (
	"errors"
	"log"
	"math/rand"
)

type Result interface {
	Error() error
	State() *State
}

type State struct {
	// Define any relevant state information
	Id      int
	Message string
	// Add more fields as needed
}

type OperationResult struct {
	err   error
	state *State
}

func (r *OperationResult) Error() error {
	return r.err
}

func (r *OperationResult) State() *State {
	return r.state
}

func SomeOperation() Result {
	// Simulate an operation
	state := &State{Id: 1, Message: "Operation completed"}

	// Simulate an error condition
	if rand.Intn(2) == 0 {
		return &OperationResult{err: errors.New("Operation failed"), state: nil}
	}

	return &OperationResult{err: nil, state: state}
}

func dispatchEvent(result Result) {
	if err := result.Error(); err != nil {
		// Handle error event
		log.Printf("Error: %v\n", err)
		dispatchErrorHandler(result)
	} else {
		// Handle success event
		state := result.State()
		log.Printf("Success: %v\n", state)
		dispatchSuccessHandler(result)
	}
}

func dispatchErrorHandler(result Result) {
	state := result.State()
	if state != nil {
		// Additional error handling based on state
		log.Printf("Error State: %v\n", state)
	}
	// Notify other systems of the error
}

func dispatchSuccessHandler(result Result) {
	state := result.State()
	// Additional success handling based on state
	log.Printf("Success State: %v\n", state)
	// Notify other systems of the success
}

func main() {
	log.Println("Starting application...")

	// Perform the operation
	result := SomeOperation()

	// Dispatch the event based on the result
	dispatchEvent(result)

	log.Println("Application finished.")
}

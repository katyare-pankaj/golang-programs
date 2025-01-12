package main

import (
	"fmt"
	"sync"
	"time"
)

// Define the states of the state machine
type State int

const (
	Idle State = iota
	Listening
	Processing
	Error
)

// Define the state machine
type StateMachine struct {
	currentState State
	waitGroup    *sync.WaitGroup
	transitions  map[State][]State
}

func (sm *StateMachine) Start() {
	sm.waitGroup = &sync.WaitGroup{}
	sm.waitGroup.Add(1)

	go sm.run()
}

func (sm *StateMachine) Stop() {
	sm.waitGroup.Done()
	sm.waitGroup.Wait()
}

func (sm *StateMachine) Transition(newState State) {
	sm.currentState = newState
	fmt.Println("State transitioned to:", sm.currentState)
}

func (sm *StateMachine) getTransitions(state State) []State {
	if transitions, ok := sm.transitions[state]; ok {
		return transitions
	}
	return nil
}

func (sm *StateMachine) run() {
	defer sm.waitGroup.Done()

	switch sm.currentState {
	case Idle:
		sm.idleState()
	case Listening:
		sm.listeningState()
	case Processing:
		sm.processingState()
	case Error:
		sm.errorState()
	}
}

func (sm *StateMachine) idleState() {
	fmt.Println("Idle state")
	time.Sleep(time.Second)

	// Simulate transitioning to Listening state
	sm.Transition(Listening)
}

func (sm *StateMachine) listeningState() {
	fmt.Println("Listening state")
	time.Sleep(time.Second)

	// Simulate transitioning to Processing state
	sm.Transition(Processing)
}

func (sm *StateMachine) processingState() {
	fmt.Println("Processing state")
	time.Sleep(time.Second)

	// Simulate transitioning to Error state
	sm.Transition(Error)
}

func (sm *StateMachine) errorState() {
	fmt.Println("Error state")
	time.Sleep(time.Second)

	// Simulate transitioning back to Idle state after handling the error
	sm.Transition(Idle)
}

func main() {
	// Define the possible transitions
	transitions := map[State][]State{
		Idle:       {Listening},
		Listening:  {Processing},
		Processing: {Error},
		Error:      {Idle},
	}

	// Create the state machine
	sm := &StateMachine{
		currentState: Idle,
		transitions:  transitions,
	}

	// Start the state machine
	sm.Start()

	// Simulate interaction with the state machine
	time.Sleep(2 * time.Second)
	sm.Stop()
}

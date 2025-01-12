package main

import (
	"fmt"
	"sync"
	"time"
)

type State int

const (
	InitialState State = iota
	Step1
	Step2
	Step3
	FinalState
)

type StateMachine struct {
	currentState State
	wg           sync.WaitGroup
}

func (sm *StateMachine) Start() {
	sm.wg.Add(1)
	defer sm.wg.Done()
	sm.transition(InitialState)
}

func (sm *StateMachine) transition(nextState State) {
	sm.currentState = nextState
	switch nextState {
	case InitialState:
		sm.handleInitialState()
	case Step1:
		sm.handleStep1()
	case Step2:
		sm.handleStep2()
	case Step3:
		sm.handleStep3()
	case FinalState:
		sm.handleFinalState()
	}
}

func (sm *StateMachine) handleInitialState() {
	fmt.Println("Entering Initial State")
	// Simulate a time-consuming task for Step 1
	time.Sleep(2 * time.Second)
	sm.transition(Step1)
}

func (sm *StateMachine) handleStep1() {
	fmt.Println("Entering Step 1")
	// Simulate a time-consuming task for Step 2
	time.Sleep(3 * time.Second)
	sm.transition(Step2)
}

func (sm *StateMachine) handleStep2() {
	fmt.Println("Entering Step 2")
	// Simulate a time-consuming task for Step 3
	time.Sleep(1 * time.Second)
	sm.transition(Step3)
}

func (sm *StateMachine) handleStep3() {
	fmt.Println("Entering Step 3")
	// Simulate a time-consuming task for Final State
	time.Sleep(4 * time.Second)
	sm.transition(FinalState)
}

func (sm *StateMachine) handleFinalState() {
	fmt.Println("Entering Final State")
	// Terminate the state machine
	sm.wg.Done()
}

func main() {
	sm := &StateMachine{}
	go sm.Start()
	fmt.Println("Waiting for the state machine to complete...")
	sm.wg.Wait()
	fmt.Println("State machine completed.")
}

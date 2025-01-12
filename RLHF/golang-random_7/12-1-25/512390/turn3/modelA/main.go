package main

import (
	"fmt"
	"sync"
	"time"
)

type State string

const (
	InitState     State = "INIT"
	Step1State    State = "STEP_1"
	Step2State    State = "STEP_2"
	Step3State    State = "STEP_3"
	CompleteState State = "COMPLETE"
)

type MultiStepProcess struct {
	currentState State
	mutex        sync.Mutex
}

func NewMultiStepProcess() *MultiStepProcess {
	return &MultiStepProcess{currentState: InitState}
}

func (m *MultiStepProcess) run() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	switch m.currentState {
	case InitState:
		fmt.Println("Starting process...")
		m.transition(Step1State)
	case Step1State:
		fmt.Println("Performing step 1...")
		m.simulateTask(Step1State)
		m.transition(Step2State)
	case Step2State:
		fmt.Println("Performing step 2...")
		m.simulateTask(Step2State)
		m.transition(Step3State)
	case Step3State:
		fmt.Println("Performing step 3...")
		m.simulateTask(Step3State)
		m.transition(CompleteState)
	case CompleteState:
		fmt.Println("Process completed.")
	default:
		fmt.Println("Invalid state:", m.currentState)
	}
}

func (m *MultiStepProcess) transition(nextState State) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.currentState = nextState
}

func (m *MultiStepProcess) simulateTask(state State) {
	fmt.Println("Simulating task for", state, "for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Task for", state, "completed.")
}

func main() {
	process := NewMultiStepProcess()
	go process.run()

	// Optionally, you can add a main loop to keep the program running
	select {}
}

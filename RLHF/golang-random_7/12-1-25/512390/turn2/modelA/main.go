package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type State string

const (
	Idle           State = "IDLE"
	Authenticating State = "AUTHENTICATING"
	Authorizing    State = "AUTHORIZING"
	AccessGranted  State = "ACCESS_GRANTED"
	AccessDenied   State = "ACCESS_DENIED"
)

type SecurityProtocol struct {
	currentState State
	wg           sync.WaitGroup
	transitions  map[State][]State
}

func NewSecurityProtocol() *SecurityProtocol {
	sp := &SecurityProtocol{
		currentState: Idle,
		wg:           sync.WaitGroup{},
		transitions: map[State][]State{
			Idle:           {Authenticating},
			Authenticating: {Authorizing, AccessDenied},
			Authorizing:    {AccessGranted},
			AccessGranted:  {Idle},
			AccessDenied:   {Idle},
		},
	}
	return sp
}

func (sp *SecurityProtocol) start() {
	sp.wg.Add(1)
	defer sp.wg.Done()
	sp.enterState(sp.currentState)
}

func (sp *SecurityProtocol) enterState(state State) {
	switch state {
	case Idle:
		sp.idle()
	case Authenticating:
		sp.authenticating()
	case Authorizing:
		sp.authorizing()
	case AccessGranted:
		sp.accessGranted()
	case AccessDenied:
		sp.accessDenied()
	default:
		fmt.Println("Invalid state:", state)
	}
}

func (sp *SecurityProtocol) transition(nextState State) {
	if !contains(sp.transitions[sp.currentState], nextState) {
		fmt.Println("Invalid transition from", sp.currentState, "to", nextState)
		return
	}

	sp.currentState = nextState
	sp.enterState(nextState)
}

func contains(s []State, e State) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (sp *SecurityProtocol) idle() {
	fmt.Println("Entering IDLE state")
	time.Sleep(time.Second * 1) // Simulate idle
	sp.transition(Authenticating)
}

func (sp *SecurityProtocol) authenticating() {
	fmt.Println("Entering AUTHENTICATING state")
	time.Sleep(time.Second * 2) // Simulate authentication
	if sp.getRandomBool() {
		sp.transition(Authorizing)
	} else {
		sp.transition(AccessDenied)
	}
}

func (sp *SecurityProtocol) authorizing() {
	fmt.Println("Entering AUTHORIZING state")
	time.Sleep(time.Second * 1) // Simulate authorization
	sp.transition(AccessGranted)
}

func (sp *SecurityProtocol) accessGranted() {
	fmt.Println("Entering ACCESS_GRANTED state")
	time.Sleep(time.Second * 3) // Simulate access session
	sp.transition(Idle)
}

func (sp *SecurityProtocol) accessDenied() {
	fmt.Println("Entering ACCESS_DENIED state")
	time.Sleep(time.Second * 1) // Simulate denied action
	sp.transition(Idle)
}

func (sp *SecurityProtocol) getRandomBool() bool {
	return sp.getRandomInt(1, 2) == 1
}

func (sp *SecurityProtocol) getRandomInt(min, max int) int {
	return min + int(rand.Intn(max-min))
}

func main() {
	securityProtocol := NewSecurityProtocol()
	go securityProtocol.start()

	fmt.Println("Waiting for security protocol to complete...")
	securityProtocol.wg.Wait()
	fmt.Println("Security protocol completed.")
}

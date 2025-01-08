package main

import (
	"fmt"
	"sync"
)

// Message is the interface for messages sent between actors.
type Message interface{}

// Actor is the interface for actors.
type Actor interface {
	Start()
	Stop()
	Send(Message)
}

// BankAccount represents a bank account actor.
type BankAccount struct {
	balance  int
	inbox    chan Message
	wg       sync.WaitGroup
	quit     chan struct{}
	stopOnce sync.Once
	mutex    sync.Mutex
}

// NewBankAccount creates a new BankAccount actor.
func NewBankAccount(balance int) *BankAccount {
	return &BankAccount{
		balance: balance,
		inbox:   make(chan Message),
		quit:    make(chan struct{}),
	}
}

// Start starts the actor's processing loop.
func (a *BankAccount) Start() {
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
func (a *BankAccount) Stop() {
	a.stopOnce.Do(func() {
		close(a.quit)
	})
}

// Send sends a message to the actor's inbox.
func (a *BankAccount) Send(msg Message) {
	a.inbox <- msg
}

// processMessage processes a message received by the actor.
func (a *BankAccount) processMessage(msg Message) {
	switch msg := msg.(type) {
	case DepositMessage:
		a.handleDepositMessage(msg)
	case WithdrawMessage:
		a.handleWithdrawMessage(msg)
	default:
		fmt.Printf("BankAccount: Unrecognized message: %v\n", msg)
	}
}

// handleDepositMessage handles a DepositMessage and updates the account balance.
func (a *BankAccount) handleDepositMessage(msg DepositMessage) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.balance += msg.Amount
	fmt.Printf("BankAccount: Deposited %d. New balance: %d\n", msg.Amount, a.balance)
}

// handleWithdrawMessage handles a WithdrawMessage and updates the account balance.
func (a *BankAccount) handleWithdrawMessage(msg WithdrawMessage) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.balance >= msg.Amount {
		a.balance -= msg.Amount
		fmt.Printf("BankAccount: Withdrew %d. New balance: %d\n", msg.Amount, a.balance)
	} else {
		fmt.Printf("BankAccount: Insufficient balance. Withdrawal of %d failed.\n", msg.Amount)
	}
}

// Wait waits for the actor to complete its processing loop.
func (a *BankAccount) Wait() {
	a.wg.Wait()
}

// DepositMessage represents a deposit message.
type DepositMessage struct {
	Amount int
}

// WithdrawMessage represents a withdrawal message.
type WithdrawMessage struct {
	Amount int
}

func main() {
	// Create a bank account actor with an initial balance of 100
	account := NewBankAccount(100)

	// Start the actor
	account.Start()

	// Send some deposit and withdrawal messages
	account.Send(DepositMessage{Amount: 50})
	account.Send(WithdrawMessage{Amount: 30})
	account.Send(WithdrawMessage{Amount: 120})
	account.Send(DepositMessage{Amount: 70})

	// Wait for the actor to complete
	account.Wait()

	// Stop the actor
	account.Stop()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// Account struct to hold localized state
type Account struct {
	balance float64
	mutex   sync.Mutex
}

// Channels for deposit and balance response
var depositChan = make(chan float64)
var balanceResponseChan = make(chan float64)

// Deposit service
func depositService(account *Account) {
	for {
		depositAmount := <-depositChan

		// Acquire lock to protect localized state
		account.mutex.Lock()
		defer account.mutex.Unlock()

		// Update balance
		account.balance += depositAmount
	}
}

// Balance service
func balanceService(account *Account) {
	for {
		// Acquire lock to protect localized state
		account.mutex.Lock()
		defer account.mutex.Unlock()

		// Read balance
		currentBalance := account.balance

		// Send balance to response channel
		balanceResponseChan <- currentBalance
	}
}

func main() {
	// Create an Account instance
	account := &Account{}

	// Start Deposit and Balance services in separate goroutines
	go depositService(account)
	go balanceService(account)

	// Simulate deposits
	go func() {
		for i := 0; i < 5; i++ {
			depositAmount := float64(i+1) * 10.0
			fmt.Printf("Depositing: %.2f\n", depositAmount)
			depositChan <- depositAmount
			time.Sleep(1 * time.Second) // simulate some delay
		}
	}()

	// Simulate balance requests
	go func() {
		for i := 0; i < 5; i++ {
			// Request balance
			currentBalance := <-balanceResponseChan
			fmt.Printf("Current Balance: %.2f\n", currentBalance)
			time.Sleep(2 * time.Second) // simulate some delay
		}
	}()

	// Sleep to allow the services to run
	time.Sleep(2 * time.Second)
	fmt.Println("Simulation finished.")
}

package main

import "sync"

// Account struct to hold localized state
type Account struct {
	balance float64
	mutex   sync.Mutex
}

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

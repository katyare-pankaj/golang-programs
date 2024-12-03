package main

import (
	"fmt"
	"sync"
)

// Account represents a bank account with a balance
type Account struct {
	mu      sync.Mutex
	balance int
}

// Deposit adds an amount to the account balance
func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

// Withdraw subtracts an amount from the account balance
func (a *Account) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance >= amount {
		a.balance -= amount
		return true
	}
	return false
}

// Transfer safely transfers money from one account to another
func Transfer(from, to *Account, amount int) bool {
	// Ensure locks are acquired in a consistent order
	// to prevent deadlocks
	if from == to {
		return false
	}

	// Ensure consistent lock order by comparing memory addresses
	if from < to {
		from.mu.Lock()
		to.mu.Lock()
	} else {
		to.mu.Lock()
		from.mu.Lock()
	}

	defer from.mu.Unlock()
	defer to.mu.Unlock()

	if from.balance >= amount {
		from.balance -= amount
		to.balance += amount
		return true
	}
	return false
}

func main() {
	accountA := &Account{balance: 1000}
	accountB := &Account{balance: 1000}

	var wg sync.WaitGroup

	// Simulate concurrent transfers
	wg.Add(2)
	go func() {
		defer wg.Done()
		if Transfer(accountA, accountB, 300) {
			fmt.Println("Transfer from A to B successful")
		} else {
			fmt.Println("Transfer from A to B failed")
		}
	}()

	go func() {
		defer wg.Done()
		if Transfer(accountB, accountA, 500) {
			fmt.Println("Transfer from B to A successful")
		} else {
			fmt.Println("Transfer from B to A failed")
		}
	}()

	// Wait for all transfers to complete
	wg.Wait()

	fmt.Printf("Final Balance A: %d, B: %d\n", accountA.balance, accountB.balance)
}

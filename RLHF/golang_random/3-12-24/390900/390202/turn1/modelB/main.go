package main

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
	mu      sync.Mutex
}

// Transaction represents a simple transaction structure
type Transaction struct {
	from   *Account
	to     *Account
	amount int
	result chan error
}

// NewAccount creates a new account with an initial balance
func NewAccount(balance int) *Account {
	return &Account{balance: balance}
}

// Transfer transfers an amount from one account to another atomically
func Transfer(tx Transaction) {
	tx.from.mu.Lock()
	defer tx.from.mu.Unlock()

	tx.to.mu.Lock()
	defer tx.to.mu.Unlock()

	if tx.from.balance < tx.amount {
		tx.result <- fmt.Errorf("transaction failed: insufficient funds")
		return
	}

	tx.from.balance -= tx.amount
	tx.to.balance += tx.amount
	tx.result <- nil
}

func main() {
	// Creating two accounts
	accountA := NewAccount(100)
	accountB := NewAccount(50)

	// Create a transaction channel
	transactionChan := make(chan Transaction)

	var wg sync.WaitGroup

	// Start transaction worker
	go func() {
		for tx := range transactionChan {
			Transfer(tx)
			wg.Done()
		}
	}()

	// Perform two transactions concurrently
	wg.Add(2)

	// Transaction 1
	go func() {
		tx := Transaction{
			from:   accountA,
			to:     accountB,
			amount: 30,
			result: make(chan error),
		}
		transactionChan <- tx
		err := <-tx.result
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Transaction 1 successful")
		}
	}()

	// Transaction 2
	go func() {
		tx := Transaction{
			from:   accountB,
			to:     accountA,
			amount: 20,
			result: make(chan error),
		}
		transactionChan <- tx
		err := <-tx.result
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Transaction 2 successful")
		}
	}()

	wg.Wait()
	close(transactionChan)

	fmt.Printf("Final balance of accountA: %d\n", accountA.balance)
	fmt.Printf("Final balance of accountB: %d\n", accountB.balance)
}

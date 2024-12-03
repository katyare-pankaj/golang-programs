package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	mu      *sync.Mutex
}

// Transaction contains the details of a transfer
type Transaction struct {
	from   *Account
	to     *Account
	amount int
}

type TransactionManager struct {
	txnChan chan Transaction
}

// NewAccount initializes a new account with a balance
func NewAccount(balance int) *Account {
	return &Account{balance: balance, mu: &sync.Mutex{}}
}

// NewTransactionManager creates a new transaction manager
func NewTransactionManager(numTxnWorkers int) *TransactionManager {
	txnChan := make(chan Transaction)
	for i := 0; i < numTxnWorkers; i++ {
		go txnWorker(txnChan)
	}
	return &TransactionManager{txnChan: txnChan}
}

// Transfer executes the transaction, ensuring it's atomic
func (manager *TransactionManager) Transfer(txn Transaction) {
	manager.txnChan <- txn
}

// txnWorker is a goroutine that processes transactions
func txnWorker(txnChan <-chan Transaction) {
	for txn := range txnChan {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Transaction error: %v\n", r)
			}
		}()

		txn.from.mu.Lock()
		txn.to.mu.Lock()
		defer func() {
			txn.from.mu.Unlock()
			txn.to.mu.Unlock()
		}()

		if txn.from.balance < txn.amount {
			panic("insufficient funds")
		}

		txn.from.balance -= txn.amount
		txn.to.balance += txn.amount
	}
}

func main() {
	// Create user accounts
	user1Account := NewAccount(100)
	user2Account := NewAccount(50)
	user3Account := NewAccount(75)

	// Set up the transaction manager
	txnManager := NewTransactionManager(3) // Use three worker goroutines

	// Simultaneously initiate transfers
	fmt.Println("Simulating concurrent transfers...")

	// Transfer 50 from user1 to user2
	go func() {
		txnManager.Transfer(Transaction{from: user1Account, to: user2Account, amount: 50})
	}()

	// Transfer 30 from user2 to user3
	go func() {
		txnManager.Transfer(Transaction{from: user2Account, to: user3Account, amount: 30})
	}()

	// Transfer 100 from user3 to user1
	go func() {
		txnManager.Transfer(Transaction{from: user3Account, to: user1Account, amount: 100})
	}()

	time.Sleep(100 * time.Millisecond)

	// Print the final account balances
	fmt.Println("Final account balances:")
	fmt.Printf("User 1: %d\n", user1Account.balance)
	fmt.Printf("User 2: %d\n", user2Account.balance)
	fmt.Printf("User 3: %d\n", user3Account.balance)
}

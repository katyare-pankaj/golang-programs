package main

import (
	"fmt"
	"sync"
)

// Account represents a bank account with a balance and a mutex for safe concurrent access.
type Account struct {
	ID      int
	Balance int
	mu      sync.Mutex
}

// Transaction represents a fund transfer from one account to another.
type Transaction struct {
	fromAccountID int
	toAccountID   int
	amount        int
	result        chan error
}

// Bank holds a collection of accounts and a transaction channel for processing transactions.
type Bank struct {
	accounts        map[int]*Account
	transactionChan chan Transaction
	wg              sync.WaitGroup
}

// NewBank creates a new bank with a given set of accounts.
func NewBank(accounts []*Account) *Bank {
	accountMap := make(map[int]*Account)
	for _, account := range accounts {
		accountMap[account.ID] = account
	}
	return &Bank{
		accounts:        accountMap,
		transactionChan: make(chan Transaction),
	}
}

// ProcessTransactions processes transactions from the transaction channel.
func (b *Bank) ProcessTransactions() {
	for tx := range b.transactionChan {
		b.wg.Add(1)
		go b.handleTransaction(tx)
	}
}

// handleTransaction handles a single transaction, ensuring atomicity.
func (b *Bank) handleTransaction(tx Transaction) {
	defer b.wg.Done()

	fromAccount := b.accounts[tx.fromAccountID]
	toAccount := b.accounts[tx.toAccountID]

	// Lock accounts in a consistent order to avoid deadlocks
	if fromAccount.ID < toAccount.ID {
		fromAccount.mu.Lock()
		toAccount.mu.Lock()
	} else {
		toAccount.mu.Lock()
		fromAccount.mu.Lock()
	}
	defer fromAccount.mu.Unlock()
	defer toAccount.mu.Unlock()

	if fromAccount.Balance < tx.amount {
		tx.result <- fmt.Errorf("transaction failed: insufficient funds in account %d", fromAccount.ID)
		return
	}

	fromAccount.Balance -= tx.amount
	toAccount.Balance += tx.amount
	tx.result <- nil
}

// Transfer initiates a transfer of funds between accounts.
func (b *Bank) Transfer(fromAccountID, toAccountID, amount int) <-chan error {
	result := make(chan error)
	tx := Transaction{
		fromAccountID: fromAccountID,
		toAccountID:   toAccountID,
		amount:        amount,
		result:        result,
	}
	b.transactionChan <- tx
	return result
}

func main() {
	// Create accounts
	accounts := []*Account{
		{ID: 1, Balance: 100},
		{ID: 2, Balance: 200},
		{ID: 3, Balance: 300},
	}

	// Create a new bank with the accounts
	bank := NewBank(accounts)

	// Start processing transactions in the bank
	go bank.ProcessTransactions()

	// Perform simultaneous transactions
	results := []<-chan error{
		bank.Transfer(1, 2, 50),
		bank.Transfer(2, 3, 70),
		bank.Transfer(3, 1, 30),
		bank.Transfer(1, 3, 80),
		bank.Transfer(2, 1, 60),
	}

	// Wait for all transactions to complete and check results
	for _, result := range results {
		if err := <-result; err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Transaction successful")
		}
	}

	// Wait for all goroutines to finish
	bank.wg.Wait()

	// Close the transaction channel
	close(bank.transactionChan)

	// Print final balances
	for _, account := range bank.accounts {
		fmt.Printf("Final balance of account %d: %d\n", account.ID, account.Balance)
	}
}

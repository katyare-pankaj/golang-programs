package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Account represents a simple bank account with a balance.
type Account struct {
	ID      int
	Balance float64
	mu      sync.Mutex // A mutex to protect concurrent access to the balance
}

// Transaction represents a financial transaction.
type Transaction struct {
	AccountID int
	Amount    float64
}

// processTransaction processes a single transaction, including error handling.
func processTransaction(account *Account, transaction Transaction) error {
	account.mu.Lock() // Ensure exclusive access to the account balance
	defer account.mu.Unlock()

	// Simulate a 10% chance of failure for transaction errors
	if rand.Intn(100) < 10 {
		return fmt.Errorf("processing failed: account ID %d transaction amount %.2f", transaction.AccountID, transaction.Amount)
	}

	newBalance := account.Balance + transaction.Amount
	if newBalance < 0 {
		return fmt.Errorf("insufficient funds for account ID %d", transaction.AccountID) // Protect against overdrawing
	}

	// Commit transaction
	account.Balance = newBalance
	fmt.Printf("Account ID %d: transaction amount %.2f succeeded, new balance: %.2f\n", transaction.AccountID, transaction.Amount, account.Balance)
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize sample accounts and transactions
	accounts := make(map[int]*Account)
	for i := 1; i <= 5; i++ {
		accounts[i] = &Account{ID: i, Balance: 100.0} // Start each with a $100 balance
	}

	transactions := []Transaction{
		{AccountID: 1, Amount: -20.0},
		{AccountID: 2, Amount: 50.0},
		{AccountID: 3, Amount: -70.0},
		{AccountID: 4, Amount: 30.0},
		{AccountID: 5, Amount: -50.0},
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	errors := make([]error, 0)

	// Process each transaction concurrently
	for _, transaction := range transactions {
		wg.Add(1) // Increment the WaitGroup counter

		go func(transaction Transaction) {
			defer wg.Done() // Decrement the counter when finished

			account, exists := accounts[transaction.AccountID]
			if !exists {
				mu.Lock()
				errors = append(errors, fmt.Errorf("unknown account ID %d", transaction.AccountID))
				mu.Unlock()
				return
			}

			// Process the transaction and capture any errors
			if err := processTransaction(account, transaction); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(transaction)
	}

	wg.Wait() // Wait for all transaction processing to complete

	// Handle and report any errors from processing
	if len(errors) > 0 {
		log.Println("Errors during transaction processing:")
		for _, err := range errors {
			log.Println(err)
		}
	} else {
		fmt.Println("All transactions processed successfully.")
	}

	// Final account balances
	fmt.Println("Final account balances:")
	for _, account := range accounts {
		fmt.Printf("Account ID %d: balance %.2f\n", account.ID, account.Balance)
	}
}

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Transaction represents a BASE transaction
type Transaction struct {
	mu        sync.Mutex
	steps     []func() error
	committed bool
	aborted   bool
}

// NewTransaction creates a new Transaction
func NewTransaction() *Transaction {
	return &Transaction{}
}

// Step adds a step to the transaction
func (t *Transaction) Step(step func() error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.committed || t.aborted {
		panic("Transaction already committed or aborted")
	}
	t.steps = append(t.steps, step)
}

// Commit commits the transaction
func (t *Transaction) Commit() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.committed || t.aborted {
		return nil
	}

	for _, step := range t.steps {
		if err := step(); err != nil {
			t.aborted = true
			return fmt.Errorf("transaction step failed: %w", err)
		}
	}

	t.committed = true
	return nil
}

// Abort aborts the transaction
func (t *Transaction) Abort() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.aborted = true
}

// IsCommitted returns true if the transaction is committed
func (t *Transaction) IsCommitted() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.committed
}

// IsAborted returns true if the transaction is aborted
func (t *Transaction) IsAborted() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.aborted
}

func main() {
	// Example usage:
	tx := NewTransaction()

	tx.Step(func() error {
		fmt.Println("Step 1: Performing some operation...")
		// Simulate an error in one of the steps
		// return fmt.Errorf("step 1 failed")
		return nil
	})

	tx.Step(func() error {
		fmt.Println("Step 2: Performing another operation...")
		time.Sleep(2 * time.Second)
		return nil
	})

	if err := tx.Commit(); err != nil {
		log.Println("Transaction failed:", err)
	} else {
		log.Println("Transaction committed successfully.")
	}
}

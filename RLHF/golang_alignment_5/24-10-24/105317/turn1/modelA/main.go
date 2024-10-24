package main

import (
	"fmt"
	"sync"
	"time"
)

// BASETransaction represents a basic transaction structure
type BASETransaction struct {
	ID        string
	Amount    int
	Status    string
	mutex     sync.Mutex
	committed bool
}

// NewBASETransaction creates a new BASETransaction
func NewBASETransaction(id string, amount int) *BASETransaction {
	return &BASETransaction{
		ID:     id,
		Amount: amount,
		Status: "Pending",
	}
}

// Commit commits the transaction
func (t *BASETransaction) Commit() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.committed {
		t.Status = "Committed"
		t.committed = true
	}
}

// Abort aborts the transaction
func (t *BASETransaction) Abort() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.committed {
		t.Status = "Aborted"
		t.committed = true
	}
}

// IsCommitted checks if the transaction is committed
func (t *BASETransaction) IsCommitted() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.committed
}

// TransactionManager manages a set of BASE transactions
type TransactionManager struct {
	transactions map[string]*BASETransaction
	mutex        sync.Mutex
}

// NewTransactionManager creates a new TransactionManager
func NewTransactionManager() *TransactionManager {
	return &TransactionManager{
		transactions: make(map[string]*BASETransaction),
	}
}

// GetTransaction retrieves a transaction by ID
func (tm *TransactionManager) GetTransaction(id string) *BASETransaction {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	return tm.transactions[id]
}

// AddTransaction adds a new transaction
func (tm *TransactionManager) AddTransaction(t *BASETransaction) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()
	tm.transactions[t.ID] = t
}

// CommitTransaction commits a transaction by ID
func (tm *TransactionManager) CommitTransaction(id string) {
	if t := tm.GetTransaction(id); t != nil {
		t.Commit()
	}
}

// AbortTransaction aborts a transaction by ID
func (tm *TransactionManager) AbortTransaction(id string) {
	if t := tm.GetTransaction(id); t != nil {
		t.Abort()
	}
}

func main() {
	// Create a new transaction manager
	tm := NewTransactionManager()

	// Create some transactions
	tx1 := NewBASETransaction("tx1", 100)
	tx2 := NewBASETransaction("tx2", 200)

	// Add transactions to the manager
	tm.AddTransaction(tx1)
	tm.AddTransaction(tx2)

	// Simulate distributed processing with goroutines
	go func() {
		// Process transaction tx1
		time.Sleep(time.Second * 2)
		tm.CommitTransaction("tx1")
		fmt.Println("Transaction tx1 committed")
	}()

	go func() {
		// Process transaction tx2
		time.Sleep(time.Second * 1)
		tm.AbortTransaction("tx2")
		fmt.Println("Transaction tx2 aborted")
	}()

	// Main thread checks transaction statuses
	for {
		time.Sleep(time.Second)
		tx1 := tm.GetTransaction("tx1")
		tx2 := tm.GetTransaction("tx2")

		fmt.Printf("Transaction tx1 status: %s\n", tx1.Status)
		fmt.Printf("Transaction tx2 status: %s\n", tx2.Status)
	}

}

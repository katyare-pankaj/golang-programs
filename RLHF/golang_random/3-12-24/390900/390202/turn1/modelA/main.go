package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Transaction struct {
	operations  []func() error
	mutex       sync.Mutex
	isCommitted bool
}

func (tx *Transaction) Commit() error {
	tx.mutex.Lock()
	defer tx.mutex.Unlock()

	if tx.isCommitted {
		return errors.New("transaction already committed")
	}

	for _, op := range tx.operations {
		err := op()
		if err != nil {
			return err
		}
	}

	tx.isCommitted = true
	return nil
}

func (tx *Transaction) Rollback() error {
	tx.mutex.Lock()
	defer tx.mutex.Unlock()

	if tx.isCommitted {
		return errors.New("transaction already committed")
	}

	for _, op := range tx.operations {
		err := op()
		if err != nil {
			return err
		}
	}

	tx.isCommitted = true
	return nil
}

func (tx *Transaction) AddOperation(op func() error) {
	tx.mutex.Lock()
	defer tx.mutex.Unlock()

	if tx.isCommitted {
		fmt.Println("transaction already committed, cannot add operation")
		return
	}

	tx.operations = append(tx.operations, op)
}

func main() {
	var err error
	tx := new(Transaction)

	// Simulate a bank account transfer operation
	tx.AddOperation(func() error {
		fmt.Println("Deducting 100 from Account A")
		return nil
	})
	tx.AddOperation(func() error {
		fmt.Println("Adding 100 to Account B")
		return nil
	})

	// Simulate concurrent access
	go func() {
		fmt.Println("Transaction starting...")

		select {
		case <-time.After(time.Second * 1):
			// Simulate an error after 1 second
			fmt.Println("An error occurred!")
			err = errors.New("transaction failed")
		default:
			fmt.Println("Transaction successful!")
			err = tx.Commit()
		}

		if err != nil {
			fmt.Println("Rolling back transaction due to error:", err)
			tx.Rollback()
		}
	}()

	time.Sleep(time.Second * 2)

	if tx.isCommitted {
		fmt.Println("Transaction committed.")
	} else {
		fmt.Println("Transaction was not committed.")
	}
}

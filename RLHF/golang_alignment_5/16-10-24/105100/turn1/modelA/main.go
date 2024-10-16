package main

import (
	"context"
	"log"
	"math/rand"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
)

const (
	databaseURL = "postgresql://root@localhost:26257/?sslmode=disable"
)

func main() {
	// Create a new database connection pool.
	db, err := crdb.Open(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer db.Close()

	// Ensure the table exists
	if _, err := db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS accounts (
			id INT PRIMARY KEY,
			balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
		)
	`); err != nil {
		log.Fatalf("error creating table: %v", err)
	}

	// Simulate a financial transaction
	for i := 0; i < 10; i++ {
		if err := transferFunds(db); err != nil {
			log.Printf("Transaction failed: %v", err)
		}
	}
}

func transferFunds(db *crdb.Conne) error {
	// Generate random account IDs
	fromAccount := rand.Intn(100) + 1
	toAccount := rand.Intn(100) + 1
	amount := rand.Float64() * 100

	// Perform the transaction in a single transaction to ensure ACID properties
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(context.Background(), `
		UPDATE accounts
		SET balance = balance - $1
		WHERE id = $2 AND balance >= $1
	`, amount, fromAccount)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `
		UPDATE accounts
		SET balance = balance + $1
		WHERE id = $2
	`, amount, toAccount)
	if err != nil {
		return err
	}

	return tx.Commit()
}

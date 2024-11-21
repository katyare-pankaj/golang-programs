package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func main() {
	// Connect to the database with a connection pool
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb?parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Set the maximum number of idle connections
	db.SetMaxIdleConns(10)
	// Set the maximum number of connections
	db.SetMaxOpenConns(100)

	// Example usage with a retry mechanism
	err = executeQueryWithRetry(db, "SELECT 1 FROM dual", 3, 5*time.Second)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
}

func executeQueryWithRetry(db *sql.DB, query string, maxAttempts int, delay time.Duration) error {
	for attempt := 0; attempt < maxAttempts; attempt++ {
		log.Printf("Attempt %d: Executing query %s", attempt+1, query)

		// Execute the query
		err := db.QueryRow(query).Scan()
		if err == nil {
			log.Printf("Query executed successfully")
			return nil
		}

		// Check if the error is recoverable
		if isRecoverableError(err) {
			log.Printf("Recoverable error: %v, retrying in %v", err, delay)
			time.Sleep(delay)
		} else {
			log.Printf("Non-recoverable error: %v", err)
			return err
		}
	}

	return errors.New(fmt.Sprintf("Failed to execute query after %d attempts", maxAttempts))
}

func isRecoverableError(err error) bool {
	// Define what constitutes a recoverable error
	// For example, network errors or database timeout errors
	// This should be tailored to your specific needs
	return errors.Is(err, nil) || errors.Is(err, sql.ErrNoRows)
}

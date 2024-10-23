package main

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func insertPerson(db *sqlx.DB, wg *sync.WaitGroup, p Person) {
	defer wg.Done()
	query := `INSERT INTO persons (first_name, last_name) VALUES (?, ?)`
	_, err := db.Exec(query, p.FirstName, p.LastName)
	if err != nil {
		log.Println("Error inserting person:", err)
	}
}

func main() {
	// Open database connection
	db, err := sqlx.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Create table if it doesn't exist
	createTable := `CREATE TABLE IF NOT EXISTS persons (
        id INT AUTO_INCREMENT PRIMARY KEY,
        first_name VARCHAR(50),
        last_name VARCHAR(50)
    )`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// Sample data to insert
	persons := []Person{
		{FirstName: "Alice", LastName: "Anderson"},
		{FirstName: "Bob", LastName: "Brown"},
		// Add more persons as needed
	}

	// Channel to buffer persons for concurrent processing
	personsChan := make(chan Person, len(persons))

	// Number of goroutines for concurrent processing
	numGoroutines := 4
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			for p := range personsChan {
				insertPerson(db, &wg, p)
			}
		}()
	}

	// Start sending persons to the channel for concurrent processing
	for _, p := range persons {
		personsChan <- p
	}

	// Close the channel to indicate that no more persons will be sent
	close(personsChan)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All persons inserted successfully!")
}

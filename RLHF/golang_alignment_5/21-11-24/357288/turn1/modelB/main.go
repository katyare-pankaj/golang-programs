package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     = "your_username"
	dbPassword = "your_password"
	dbName     = "your_database_name"
	numWorkers = 100
	queueSize  = 1000
)

type job struct {
	// Your job data here
	data string
}

func main() {
	// Initialize the database connection pool
	db, err := initDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Start the worker pool
	jobQueue := make(chan job, queueSize)
	wg := sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, db, jobQueue)
	}

	// Create a HTTP server to handle incoming requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Receive new jobs and put them in the queue
		job := job{
			// Your job data here
			data: r.URL.Path,
		}
		select {
		case jobQueue <- job:
			fmt.Fprintln(w, "Job added to queue")
		default:
			fmt.Fprintln(w, "Queue is full, try again later")
		}
	})

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Wait for all workers to finish
	wg.Wait()
}

func worker(wg *sync.WaitGroup, db *sql.DB, jobQueue <-chan job) {
	defer wg.Done()

	for job := range jobQueue {
		processJob(db, job)
	}
}

func processJob(db *sql.DB, job job) {
	// Simulate some work with a random delay
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Perform the database operation
	if err := executeDBQuery(db, job); err != nil {
		// Log the error and attempt recovery
		log.Println("Database error:", err)
		recoverDB(db)
	}
}

func initDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(numWorkers)     // Set maximum number of open connections to the database
	db.SetConnMaxLifetime(time.Minute) // Set connection maximum lifetime
	return db, nil
}

func executeDBQuery(db *sql.DB, job job) error {
	// Example database query
	query := "INSERT INTO jobs (data) VALUES (?)"
	_, err := db.ExecContext(context.Background(), query, job.data)
	return err
}

func recoverDB(db *sql.DB) {
	// Example recovery mechanism: close the current connection and reopen it
	db.Close()
	db, err := initDB()
	if err != nil {
		log.Fatal("Failed to recover database:", err)
	}
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

const (
	numWorkers = 10
)

var (
	wg sync.WaitGroup
)

// Function to handle worker threads
func worker(id int, ch chan net.Conn) {
	defer wg.Done()
	for conn := range ch {
		// Perform network security checks on the connection
		// Simulate some security check by sleeping for a random duration
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		// Process the connection (e.g., read/write data, close the connection)
		conn.Close()
	}
}

func main() {
	// Create a channel to dispatch connections to workers
	connCh := make(chan net.Conn, 100)

	// Start the worker threads
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, connCh)
	}

	// Listen on a TCP port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening on port: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server started on port 8080")

	// Handle incoming connections and dispatch them to workers
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Dispatch the connection to a worker thread
		connCh <- conn
	}
}

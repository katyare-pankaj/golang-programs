package main

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

const (
	numConnections = 10
	workDuration   = 5 * time.Second
)

func main() {
	// Initialize the random seed
	rand.Seed(time.Now().UnixNano())

	// Create a wait group to ensure all connections are handled before exit
	var wg sync.WaitGroup
	wg.Add(numConnections)

	// Create a channel to signal when connections are closed
	closeCh := make(chan struct{})

	// Placeholder event loop that simulates handling of connections
	for i := 0; i < numConnections; i++ {
		go func() {
			defer func() {
				wg.Done()
				close(closeCh)
			}()

			// Simulate handling a new connection (substitute with real connection in a practical scenario)
			conn := simulateNewConnection()
			if conn == nil {
				return
			}
			defer conn.Close()

			// Handle the connection
			handleConnection(conn)
		}()
	}

	// Wait for all connections to be handled
	wg.Wait()

	// Close all channels and ensure any pending work completes
	close(closeCh)

	fmt.Println("All connections handled. Exiting...")
}

// simulateNewConnection simulates a new network connection. In practice, this could be a listener accepting a connection.
func simulateNewConnection() net.Conn {
	// Here we're using a udp 'dummy' connection simulating an incoming connection
	conn, err := net.Dial("udp", "localhost:0")
	if err != nil {
		fmt.Println("Error simulating connection:", err)
		return nil
	}
	fmt.Println("Simulated new connection.")
	return conn
}

// handleConnection is a placeholder function representing the work done on a connection
func handleConnection(conn net.Conn) {
	fmt.Println("Handling connection:", conn.LocalAddr().String())
	time.Sleep(workDuration)
	fmt.Println("Done handling:", conn.LocalAddr().String())
}

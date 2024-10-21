package main

import (
	"log"
	"net"
	"time"
)

// handleClient handles a single client connection
func handleClient(conn net.Conn) {
	defer conn.Close()
	// Buffer for reading data
	buf := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Error reading from client: %v", err)
			return
		}

		// Process the received data
		data := buf[:n]
		response := processData(data)

		// Write the response back to the client
		_, err = conn.Write(response)
		if err != nil {
			log.Printf("Error writing to client: %v", err)
			return
		}
	}
}

// processData simulates data processing and generates a response
func processData(data []byte) []byte {
	// Simulate some work
	time.Sleep(50 * time.Millisecond)
	return append(data, []byte(" - Processed")...)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()

	log.Println("Server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleClient(conn) // Handle client connections concurrently
	}
}

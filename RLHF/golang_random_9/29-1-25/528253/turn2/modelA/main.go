package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Start a TCP server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server started on :8080")

	// Infinite loop to accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected to:", conn.RemoteAddr())

	// Use a buffered reader to read lines from the connection
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		// Process message (e.g., log it)
		fmt.Printf("Message received from %s: %s", conn.RemoteAddr(), message)

		// Echo message back to client
		_, err = conn.Write([]byte("Message received: " + message))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

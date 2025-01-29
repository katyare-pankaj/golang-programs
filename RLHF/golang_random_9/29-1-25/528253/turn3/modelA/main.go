package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from:", conn.RemoteAddr())

	// Set a deadline for reading. Read operation will fail after the timeout.
	conn.SetReadDeadline(time.Now().Add(5 * time.Minute))

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Printf("Received message from %s: %s", conn.RemoteAddr(), message)

		_, err = conn.Write([]byte("Message received\n"))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

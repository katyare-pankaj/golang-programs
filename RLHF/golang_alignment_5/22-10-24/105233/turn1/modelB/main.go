package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a Unix domain socket listener
	listener, err := net.Listen("unix", "/tmp/microservice.sock")
	if err != nil {
		fmt.Println("Error listening on socket:", err)
		return
	}
	defer listener.Close()

	// Accept incoming connections
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	// Create a buffer for receiving data
	buf := make([]byte, 4096)

	// Perform zero-copy I/O
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading data:", err)
			break
		}

		// Process the received data (e.g., echo it back to the client)
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Error writing data:", err)
			break
		}
	}
}

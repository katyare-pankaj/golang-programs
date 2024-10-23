package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	bufferSize = 32 * 1024 // 32KB buffer size
)

// zeroCopyStreamHandler handles a single streaming connection using zero-copy techniques.
func zeroCopyStreamHandler(conn net.Conn) {
	defer conn.Close()

	// Open the file to be streamed
	file, err := os.Open("example.txt")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	// Get the file info to set the Content-Length header
	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Error stat file: %v", err)
		return
	}

	// Write the HTTP header with Content-Length
	_, err = conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\n", fileInfo.Size())))
	if err != nil {
		log.Printf("Error writing header: %v", err)
		return
	}

	// Create a buffer for zero-copy reads
	buffer := make([]byte, bufferSize)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			// End of file, break the loop
			break
		}
		if err != nil {
			log.Printf("Error reading file: %v", err)
			return
		}

		// Perform zero-copy write to the connection
		_, err = conn.Write(buffer[:n])
		if err != nil {
			log.Printf("Error writing to connection: %v", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle each connection in a goroutine for concurrency
		go zeroCopyStreamHandler(conn)
	}
}

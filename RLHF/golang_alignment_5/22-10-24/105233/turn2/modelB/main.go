package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func main() {
	// Replace this with the actual file path
	filePath := "/path/to/the/large/file"

	// Create a TCP connection to the FileReceiver microservice
	conn, err := net.Dial("tcp", "fileReceiver:12345")
	if err != nil {
		fmt.Println("Error connecting to FileReceiver:", err)
		return
	}
	defer conn.Close()

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Perform zero-copy I/O
	info, err := file.Stat()
	if err != nil {
		fmt.Println("Error stat file:", err)
		return
	}

	fmt.Printf("Transferring file %s of size %d bytes\n", filePath, info.Size())

	// Get the file descriptor of the TCP connection
	connFile, err := conn.(*net.TCPConn).File()
	if err != nil {
		fmt.Println("Error getting connection file descriptor:", err)
		return
	}
	defer connFile.Close() // Don't forget to close this file descriptor

	// Cast file size to int (assuming file size is within int range)
	fileSize := int(info.Size())

	// Use Sendfile to transfer the file
	_, err = syscall.Sendfile(int(connFile.Fd()), int(file.Fd()), nil, fileSize)
	if err != nil {
		fmt.Println("Error transferring file:", err)
		return
	}

	fmt.Println("File transfer completed successfully.")
}

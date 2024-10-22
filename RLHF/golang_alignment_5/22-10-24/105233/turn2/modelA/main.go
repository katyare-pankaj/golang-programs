package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"syscall"
)

const (
	bufferSize = 1024 * 1024 // 1MB buffer size
)

func main() {
	// sender
	sendFile := "example.txt"
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	f, err := os.Open(sendFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Get the file information
	fi, err := f.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	// Send the file size
	err = conn.Write([]byte(fmt.Sprintf("%d", fi.Size())))
	if err != nil {
		fmt.Println("Error sending file size:", err)
		return
	}

	// Perform zero-copy file transfer from sender to receiver
	sendBytes, err := zeroCopySendFile(conn, f)
	if err != nil {
		fmt.Println("Error sending file:", err)
		return
	}

	fmt.Printf("Sent %d bytes successfully\n", sendBytes)

	// receiver
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	conn, err = listener.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err)
		return
	}
	defer conn.Close()

	// Receive the file size
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading file size:", err)
		return
	}

	fileSize := int64(atoi(string(buf[:n])))
	fmt.Printf("Receiving file of size: %d bytes\n", fileSize)

	// Create a new file to store the received data
	receiveFile := "received_file.txt"
	f, err = os.Create(receiveFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	// Perform zero-copy file transfer from receiver to local file
	receivedBytes, err := zeroCopyReceiveFile(conn, f, fileSize)
	if err != nil {
		fmt.Println("Error receiving file:", err)
		return
	}

	fmt.Printf("Received %d bytes successfully and saved as %s\n", receivedBytes, receiveFile)
}

func zeroCopySendFile(conn net.Conn, file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	mfd, err := syscall.MemfdCreate("zerocopy_send", 0)
	if err != nil {
		return 0, err
	}
	defer syscall.Close(mfd)

	err = syscall.Ftruncate(mfd, fileInfo.Size())
	if err != nil {
		return 0, err
	}

	mapAddr, err := syscall.Mmap(mfd, 0, int(fileInfo.Size()), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return 0, err
	}
	defer syscall.Munmap(mapAddr, int(fileInfo.Size()))

	n, err := file.Read(mapAddr)
	if err != nil && err != io.EOF {
		return 0, err
	}

	sentBytes, err := conn.Write(mapAddr)
	if err != nil {
		return 0, err
	}

	return int64(sentBytes), nil
}

func zeroCopyReceiveFile(conn net.Conn, file *os.File, fileSize int64) (int64, error) {
	mfd, err := syscall.MemfdCreate("zerocopy_recv", 0)
	if err != nil {
		return 0, err
	}
	defer syscall.Close(mfd)

	err = syscall.Ftruncate(mfd, fileSize)
	if err != nil {
		return 0, err
	}

	mapAddr, err := syscall.Mmap(mfd, 0, int(fileSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return 0, err
	}
	defer syscall.Munmap(mapAddr, int(fileSize))

	receivedBytes, err := conn.Read(mapAddr)
	if err != nil && err != io.EOF {
		return 0, err
	}

	n, err := file.Write(mapAddr[:receivedBytes])
	if err != nil {
		return 0, err
	}

	return int64(receivedBytes), nil
}

func atoi(str string) int {
	num := 0
	for _, ch := range str {
		if ch < '0' || ch > '9' {
			return num
		}
		num = num*10 + int(ch-'0')
	}
	return num
}

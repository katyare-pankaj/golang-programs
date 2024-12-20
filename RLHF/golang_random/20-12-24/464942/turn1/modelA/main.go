package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after the function returns

	// Start a Goroutine to read the file
	go func() {
		buf := make([]byte, 1024)
		n, err := file.Read(buf)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("Read", n, "bytes from file:", string(buf[:n]))
		time.Sleep(1 * time.Second) // Simulate some work
	}()

	// Do other work in the main Goroutine
	fmt.Println("Main Goroutine doing other work...")
	time.Sleep(2 * time.Second)
}

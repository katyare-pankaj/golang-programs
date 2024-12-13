package main

import (
	"fmt"
	"os"
	"sync"
)

// Function that simulates a worker process
func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Processed:", num)
	}
}

func main() {
	var wg sync.WaitGroup

	// Create a channel
	ch := make(chan int)
	// Ensure the channel is closed when main function exits
	defer close(ch) // this defer statement makes sure the channel will be closed later

	wg.Add(1)          // Add one to the wait group for the worker
	go worker(ch, &wg) // Start the worker goroutine

	// Sending data to the channel
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// Here we naturally don't need to close the channel manually already since it's deferred above.

	// Wait for the worker to finish processing
	wg.Wait()
	fmt.Println("All processing done.")

	// Opening a file for demonstration purpose
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file is closed when main function exits
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Write to the file
	_, err = file.WriteString("Hello, World!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and content written.")
}

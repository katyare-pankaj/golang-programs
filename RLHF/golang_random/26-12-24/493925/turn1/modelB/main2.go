package main

import (
	"fmt"
	"time"
)

func worker(id int, done chan bool) {
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
	done <- true // Signal completion
}

func main() {
	numWorkers := 5
	done := make(chan bool, numWorkers) // Buffered channel to avoid blocking

	for i := 0; i < numWorkers; i++ {
		go worker(i, done)
	}

	// Receive completion signals from workers
	for i := 0; i < numWorkers; i++ {
		<-done
	}
	fmt.Println("All workers finished.")
}

package main

import (
	"fmt"
	"time"
)

func worker(id int, out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- id * i
		time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	}
}

func main() {
	// Create send-only channels for each worker
	worker1Out := make(chan int)
	worker2Out := make(chan int)

	// Start workers
	go worker(1, worker1Out)
	go worker(2, worker2Out)

	// Receive data from both workers using a select statement
	for {
		select {
		case value, open := <-worker1Out:
			if !open {
				fmt.Println("Worker 1 channel closed.")
				break
			}
			fmt.Println("Received from worker 1:", value)
		case value, open := <-worker2Out:
			if !open {
				fmt.Println("Worker 2 channel closed.")
				break
			}
			fmt.Println("Received from worker 2:", value)
		default:
			fmt.Println("Waiting for data...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

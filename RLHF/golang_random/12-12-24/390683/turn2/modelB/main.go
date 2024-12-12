package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, send chan<- int, receive <-chan int) {
	wg.Add(1)
	defer wg.Done()

	select {
	case send <- id:
		fmt.Printf("Worker %d sent message\n", id)
		// Simulate work
		time.Sleep(1 * time.Second)

	case <-time.After(5 * time.Second):
		fmt.Printf("Worker %d timed out sending message\n", id)
	}

	select {
	case msg := <-receive:
		fmt.Printf("Worker %d received message: %d\n", id, msg)
	case <-time.After(5 * time.Second):
		fmt.Printf("Worker %d timed out receiving message\n", id)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go worker(1, ch1, ch2)
	go worker(2, ch2, ch1)

	wg.Wait()
}

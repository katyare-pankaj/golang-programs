package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ch chan int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	for {
		select {
		case msg := <-ch:
			fmt.Printf("Worker %d received message: %d\n", id, msg)
			time.Sleep(time.Second)
		default:
			fmt.Printf("Worker %d checking for messages...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	const numWorkers = 3
	ch := make(chan int)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("Sent message: %d\n", i)
			time.Sleep(time.Second)
		}
		close(ch) // Signal the end of the channel
	}()

	wg.Wait()
}

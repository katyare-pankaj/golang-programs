package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, ch chan bool) {
	defer wg.Done()
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d is done.\n", id)
	ch <- true
}

func main() {
	var wg sync.WaitGroup
	// Buffered channel to handle completion signals
	ch := make(chan bool, 5)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, i, ch)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Iterate over the buffered channel to ensure all workers are done
	close(ch) // Close the channel to signal all tasks are processed
	for range ch {
	}

	fmt.Println("All workers have completed successfully.")
}

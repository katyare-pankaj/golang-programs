package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, ch chan bool) {
	defer wg.Done()
	// Introduce random sleep time to simulate varying work load
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Printf("Worker %d is done.\n", id)
	ch <- true // Signal completion to the main goroutine
}

func main() {
	const numWorkers = 10
	var wg sync.WaitGroup
	ch := make(chan bool, numWorkers) // Buffered channel with buffer size equal to the number of workers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i, ch)
	}

	// Collect completion signals from all workers
	for i := 0; i < numWorkers; i++ {
		<-ch
		fmt.Printf("Worker %d was processed.\n", i)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers are done.")
}

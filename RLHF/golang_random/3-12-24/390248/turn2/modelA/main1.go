package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu         sync.Mutex
	cond       = sync.NewCond(&mu)
	workQueue  []int
	workDone   int
	done       bool
	numWorkers int
)

func worker() {
	for !done {
		mu.Lock()
		for len(workQueue) == 0 {
			cond.Wait()
		}
		work := workQueue[0]
		workQueue = workQueue[1:]
		workDone++
		mu.Unlock()
		fmt.Println("Processing work:", work)
		time.Sleep(1 * time.Second) // Simulate work
	}
}

func main() {
	numWorkers = 1
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	// Add work items
	for i := 0; i < 5; i++ {
		mu.Lock()
		workQueue = append(workQueue, i)
		cond.Signal()
		mu.Unlock()
	}

	// Wait for all work to be done
	mu.Lock()
	for workDone < 5 {
		cond.Wait()
	}
	done = true
	cond.Broadcast()
	mu.Unlock()
}

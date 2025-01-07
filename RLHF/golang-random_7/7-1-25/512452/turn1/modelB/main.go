package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 10
	numEvents  = 1_000_000
	maxDelay   = 10 * time.Millisecond
)

type event struct {
	ID int
}

func worker(wg *sync.WaitGroup, eventChannel chan event) {
	defer wg.Done()
	for e := range eventChannel {
		processEvent(e)
	}
}

func processEvent(e event) {
	// Simulate random processing time
	time.Sleep(time.Duration(rand.Int63n(int64(maxDelay))))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	eventChannel := make(chan event, numEvents)

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, eventChannel)
	}

	// Generate and send events
	startTime := time.Now()
	for i := 0; i < numEvents; i++ {
		eventChannel <- event{ID: i}
	}

	// Wait for all workers to finish
	wg.Wait()
	close(eventChannel)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("Elapsed time: %s\n", elapsedTime)
	fmt.Printf("Throughput: %.2f events/second\n", float64(numEvents)/elapsedTime.Seconds())
}

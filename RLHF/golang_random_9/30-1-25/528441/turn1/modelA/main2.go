package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 5

func worker(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range data {
		fmt.Printf("Worker %d processing data: %d\n", id, v)
		// Simulate processing time
		time.Sleep(500 * time.Millisecond)
	}
}

func runWorkers(dataStream <-chan int) {
	dataCh := distributeLoad(dataStream, numWorkers)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataCh[i], &wg)
	}
	wg.Wait()
}

// Function to distribute load in a round-robin manner
func distributeLoad(data <-chan int, numWorkers int) []chan int {
	out := make([]chan int, numWorkers)
	for i := range out {
		out[i] = make(chan int, 10) // Buffer size can be adjustable
	}

	go func() {
		defer func() {
			for _, ch := range out {
				close(ch)
			}
		}()
		for v := range data {
			selected := v % numWorkers
			out[selected] <- v
		}
	}()

	return out
}

func main() {
	// ... assuming dataStream is created as shown previously
	runWorkers(dataStream)
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type monitorData struct {
	sync.Mutex // Mutex to protect concurrent access
	counter    int
}

func (md *monitorData) increment() {
	md.Lock()
	defer md.Unlock()
	md.counter++
}

func (md *monitorData) getCounter() int {
	md.Lock()
	defer md.Unlock()
	return md.counter
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup, md *monitorData) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d received timeout, exiting.\n", id)
			return
		default:
			time.Sleep(time.Millisecond * 100) // Simulate work
			md.increment()
			fmt.Printf("Worker %d incremented counter to %d\n", id, md.getCounter())
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var md monitorData

	numWorkers := 3
	timeout := 1 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // Ensure the context is cancelled to clean up resources

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg, &md)
	}

	wg.Wait()

	fmt.Printf("Final counter value: %d\n", md.getCounter())
}

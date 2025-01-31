package main

import (
	"fmt"
	"sync"
	"time"
)

type monitorData struct {
	sync.Mutex
	counter int
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

func worker(id int, wg *sync.WaitGroup, md *monitorData, done chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Printf("Worker %d stopped.\n", id)
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
	done := make(chan struct{})
	timeoutDuration := 5 * time.Second // Set your desired timeout duration

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, &md, done)
	}

	go func() {
		time.Sleep(timeoutDuration)
		fmt.Println("Timeout reached. Stopping workers...")
		close(done) // Close the done channel to signal workers to stop
	}()

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", md.getCounter())
}

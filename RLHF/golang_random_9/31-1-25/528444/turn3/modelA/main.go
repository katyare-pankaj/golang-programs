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

func worker(id int, wg *sync.WaitGroup, md *monitorData, done, pause, resume chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Printf("Worker %d stopped.\n", id)
			return
		case <-pause:
			fmt.Printf("Worker %d paused.\n", id)
			<-resume // Wait until resume signal is received
			fmt.Printf("Worker %d resumed.\n", id)
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
	pause := make(chan struct{})
	resume := make(chan struct{})

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, &md, done, pause, resume)
	}

	// Simulate pausing and resuming
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Pausing workers...")
		close(pause) // Signal workers to pause

		time.Sleep(2 * time.Second)
		fmt.Println("Resuming workers...")
		close(resume) // Signal workers to resume
	}()

	// Simulate stopping after some time
	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("Stopping workers...")
		close(done) // Signal workers to stop
	}()

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", md.getCounter())
}

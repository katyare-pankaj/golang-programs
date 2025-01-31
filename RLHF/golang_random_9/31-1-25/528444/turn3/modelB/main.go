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

func worker(id int, wg *sync.WaitGroup, md *monitorData, pause chan struct{}, resume chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-pause:
			fmt.Printf("Worker %d paused.\n", id)
		case <-resume:
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
	pause := make(chan struct{})
	resume := make(chan struct{})

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, &md, pause, resume)
	}

	time.Sleep(2 * time.Second) // Simulate some work

	fmt.Println("Pausing workers...")
	close(pause)

	time.Sleep(2 * time.Second) // Simulate pause duration

	fmt.Println("Resuming workers...")
	close(resume)

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", md.getCounter())
}

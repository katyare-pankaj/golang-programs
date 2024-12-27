package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d started on core %d\n", id, runtime.GOMAXPROCS(0))
	time.Sleep(time.Duration(id) * time.Second) // Delay based on worker ID for more visual diversity
	fmt.Printf("Worker %d finished on core %d\n", id, runtime.GOMAXPROCS(0))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores
	fmt.Printf("Number of available CPU cores: %d\n", runtime.NumCPU())

	var wg sync.WaitGroup
	const numWorkers = 20
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(&wg, i)
	}

	fmt.Println("Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("All workers finished.")
}

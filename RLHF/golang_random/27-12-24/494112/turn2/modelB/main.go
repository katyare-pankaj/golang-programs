package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d started on core %d\n", id, runtime.NumCPU())
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d finished on core %d\n", id, runtime.NumCPU())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores
	fmt.Printf("Number of available CPU cores: %d\n", runtime.NumCPU())
	var wg sync.WaitGroup
	const numWorkers = 10
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, i)
	}
	fmt.Println("Waiting for workers to finish...")
	wg.Wait()
	fmt.Println("All workers finished.")
}

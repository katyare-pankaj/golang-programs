package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		// Simulate work by performing some calculations
		_ = i * i * i
	}
	fmt.Printf("Worker %d finished on core %d\n", id, runtime.NumCPU())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores
	fmt.Printf("Number of available CPU cores: %d\n", runtime.NumCPU())
	var wg sync.WaitGroup
	const numWorkers = 10
	wg.Add(numWorkers)

	start := time.Now()
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, i)
	}
	fmt.Println("Waiting for workers to finish...")
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("All workers finished in %s\n", elapsed)
}
